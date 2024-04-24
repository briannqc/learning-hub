package app

import (
	"time"

	"go.temporal.io/sdk/workflow"
)

// GreetingWorkflow - Workflows are functions that define the overall flow of the application
// and represent the orchestration aspect of the business logic.
//
// You can pass multiple inputs to a Workflow, but it's a good practice to send a single input.
// If you have several values you want to send, you should define a Struct and pass that into
// the Workflow instead.
func GreetingWorkflow(ctx workflow.Context, name string) (string, error) {
	options := workflow.ActivityOptions{
		StartToCloseTimeout: 5 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	var result string
	err := workflow.ExecuteActivity(ctx, ComposeGreeting, name).Get(ctx, &result)
	return result, err
}
