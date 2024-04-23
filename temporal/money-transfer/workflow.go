package app

import (
	"fmt"
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

func MakeTransfer(ctx workflow.Context, input PaymentDetails) (string, error) {
	// RetryPolicy specifies how to automatically handle retries if an Activity fails
	retryPolicy := &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    time.Minute,
		MaximumAttempts:    10,
		NonRetryableErrorTypes: []string{
			"InvalidAccountError",
			"InsufficientFundsError",
		},
	}

	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy:         retryPolicy,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	var withdrawOutput string
	if withdrawErr := workflow.ExecuteActivity(ctx, Withdraw, input).Get(ctx, &withdrawOutput); withdrawErr != nil {
		return "", withdrawErr
	}

	var depositOutput string
	depositErr := workflow.ExecuteActivity(ctx, Deposit, input).Get(ctx, &depositOutput)
	if depositErr == nil {
		result := fmt.Sprintf("Transfer complete (transaction IDs: %s, %s)", withdrawOutput, depositOutput)
		return result, nil
	}

	var refundResult string
	refundErr := workflow.ExecuteActivity(ctx, Refund, input).Get(ctx, &refundResult)
	if refundErr != nil {
		return "", fmt.Errorf("failed to deposit money to: %v: %v, and money could not be returned to: %v: %w",
			input.TargetAccount, depositErr, input.SourceAccount, refundErr)
	}
	return "", fmt.Errorf("failed to deposit money to: %v: %w, and money returned to: %v",
		input.TargetAccount, depositErr, input.SourceAccount)
}
