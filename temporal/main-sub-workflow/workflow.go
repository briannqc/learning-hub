package app

import (
	"context"
	"time"

	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

type OnboardInput struct {
	Name string
	ID   string
	Date time.Time
}

type OnboardOutput struct {
	Successful bool
}

func OnboardWorkflow(ctx workflow.Context, input OnboardInput) (OnboardOutput, error) {
	logger := workflow.GetLogger(ctx)
	logger.Info("Onboard workflow started")

	hrOnboardFuture := workflow.ExecuteChildWorkflow(ctx, HROnboardWorkflow, input)
	itOnboardFuture := workflow.ExecuteChildWorkflow(ctx, ITOnboardWorkflow, input)

	var hrOnboardOutput OnboardOutput
	if err := hrOnboardFuture.Get(ctx, &hrOnboardOutput); err != nil {
		logger.Error("HR onboarding failed.", "Error", err)
		return OnboardOutput{}, err
	}

	var itOnboardOutput OnboardOutput
	if err := itOnboardFuture.Get(ctx, &itOnboardOutput); err != nil {
		logger.Error("IT onboarding failed.", "Error", err)
		return OnboardOutput{}, err
	}

	return OnboardOutput{Successful: true}, nil
}

func HROnboardWorkflow(ctx workflow.Context, input OnboardInput) (OnboardOutput, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: 5 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	logger := workflow.GetLogger(ctx)
	logger.Info("HROnboard workflow started")

	var output OnboardOutput
	if err := workflow.ExecuteActivity(ctx, SignEmploymentAgreement, input).Get(ctx, &output); err != nil {
		logger.Info("HROnboard workflow failed due to sign employment agreement.", "Error", err)
		return output, err
	}

	if err := workflow.ExecuteActivity(ctx, SignNonDisclosureAgreement, input).Get(ctx, &output); err != nil {
		logger.Info("HROnboard workflow failed due to sign non disclosure agreement.", "Error", err)
		return output, err
	}

	logger.Info("HROnboard workflow completed")
	return OnboardOutput{Successful: true}, nil
}

func SignEmploymentAgreement(ctx context.Context, input OnboardInput) (OnboardOutput, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("SignEmploymentAgreement completed")
	return OnboardOutput{Successful: true}, nil
}

func SignNonDisclosureAgreement(ctx context.Context, input OnboardInput) (OnboardOutput, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("SignNonDisclosureAgreement completed")
	return OnboardOutput{Successful: true}, nil
}

func ITOnboardWorkflow(ctx workflow.Context, input OnboardInput) (OnboardOutput, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: 5 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, options)
	logger := workflow.GetLogger(ctx)
	logger.Info("ITOnboardWorkflow started")

	var output OnboardOutput
	if err := workflow.ExecuteActivity(ctx, ActivateAccount, input).Get(ctx, &output); err != nil {
		logger.Info("ITOnboard workflow failed due to active account.", "Error", err)
		return output, err
	}

	if err := workflow.ExecuteActivity(ctx, IssueLaptop, input).Get(ctx, &output); err != nil {
		logger.Info("ITOnboard workflow failed due to issue laptop.", "Error", err)
		return output, err
	}

	if err := workflow.ExecuteActivity(ctx, IssueMonitor, input).Get(ctx, &output); err != nil {
		logger.Info("ITOnboard workflow failed due to issue monitor.", "Error", err)
		return output, err
	}

	logger.Info("ITOnboard workflow completed")
	return OnboardOutput{Successful: true}, nil
}

func ActivateAccount(ctx context.Context, input OnboardInput) (OnboardOutput, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("ActivateAccount completed")
	return OnboardOutput{Successful: true}, nil
}

func IssueLaptop(ctx context.Context, input OnboardInput) (OnboardOutput, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("IssueLaptop completed")
	return OnboardOutput{Successful: true}, nil
}

func IssueMonitor(ctx context.Context, input OnboardInput) (OnboardOutput, error) {
	logger := activity.GetLogger(ctx)
	logger.Info("IssueMonitor completed")
	return OnboardOutput{Successful: true}, nil
}
