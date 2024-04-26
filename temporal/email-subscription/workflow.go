package emailsub

import (
	"errors"
	"time"

	"go.temporal.io/sdk/workflow"
)

func SubscriptionWorkflow(ctx workflow.Context, details EmailDetails) error {
	duration := 12 * time.Second
	logger := workflow.GetLogger(ctx)
	logger.Info("Subscription created", "EmailAddress", details.EmailAddress)
	err := workflow.SetQueryHandler(ctx, "GetDetails", func() (EmailDetails, error) {
		return details, nil
	})
	if err != nil {
		return err
	}
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 2 * time.Minute,
		WaitForCancellation: true,
	})
	defer cleanUp(ctx, details)

	logger.Info("Sending welcome email", "EmailAddress", details.EmailAddress)
	details.SubscriptionCount++
	data := EmailDetails{
		EmailAddress:      details.EmailAddress,
		Message:           "Welcome! Looks like you've been signed up!",
		IsSubscribed:      true,
		SubscriptionCount: details.SubscriptionCount,
	}

	if err := workflow.ExecuteActivity(ctx, SendEmail, data).Get(ctx, nil); err != nil {
		return err
	}

	for details.IsSubscribed {
		details.SubscriptionCount++
		data := EmailDetails{
			EmailAddress:      details.EmailAddress,
			Message:           "This is yet another email in the subscription Workflow.",
			IsSubscribed:      true,
			SubscriptionCount: details.SubscriptionCount,
		}

		if err := workflow.ExecuteActivity(ctx, SendEmail, data).Get(ctx, nil); err != nil {
			return err
		}

		logger.Info("Sent content email", "EmailAddress", details.EmailAddress)
		if err := workflow.Sleep(ctx, duration); err != nil {
			return err
		}
	}
	return nil
}

func cleanUp(ctx workflow.Context, details EmailDetails) {
	logger := workflow.GetLogger(ctx)
	newCtx, cancel := workflow.NewDisconnectedContext(ctx)
	defer cancel()

	if errors.Is(ctx.Err(), workflow.ErrCanceled) {
		data := EmailDetails{
			EmailAddress:      details.EmailAddress,
			Message:           "Your subscription has been canceled. Sorry to see you go!",
			IsSubscribed:      false,
			SubscriptionCount: details.SubscriptionCount,
		}
		if err := workflow.ExecuteActivity(newCtx, SendEmail, data).Get(newCtx, nil); err != nil {
			logger.Error("Failed to send cancellation email", "Error", err)
		} else {
			logger.Info("Sent cancellation email", "EmailAddress", details.EmailAddress)
		}
	}
}
