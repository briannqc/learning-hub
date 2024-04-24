package app

import (
	"context"
	"fmt"
)

// ComposeGreeting - Activities are functions called during Workflow Execution and represent the
// execution aspect of your business logic.
//
// Activities are where you execute non-deterministic code or perform operations that may fail,
// such as API requests or database calls. Your Workflow Definitions call Activities and process
// the results. Complex Temporal Applications have Workflows that invoke many Activities, using
// the results of one Activity to execute another.
//
// Like Workflow Definitions, if you have more than one parameter for an Activity, you should
// bundle the data into a struct rather than sending multiple input parameters. This will make
// future updates easier.
func ComposeGreeting(ctx context.Context, name string) (string, error) {
	greeting := fmt.Sprintf("Hello, %s!", name)
	return greeting, nil
}
