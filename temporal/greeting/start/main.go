package main

import (
	"context"
	"fmt"
	"log"

	"go.temporal.io/sdk/client"

	app "github.com/briannqc/learning-hub/temporal/greeting"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to connect to temporal server", err)
	}
	defer c.Close()

	// You don't need to specify a Workflow ID, as Temporal will generate one for you, but
	// defining the Workflow ID yourself makes it easier for you to find it later in logs or
	// interact with a running Workflow in the future.
	//
	// Using a Workflow ID that reflects some business process or entity is a good practice.
	// For example, you might use a customer identifier or email address as part of the
	// Workflow ID if you ran one Workflow per customer. This would make it easier to find
	// all the Workflow Executions related to that customer later.
	workflowID := "hello-world-temporal"
	options := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: app.GreetingTaskQueue,
	}

	name := "Brian"
	we, err := c.ExecuteWorkflow(context.Background(), options, app.GreetingWorkflow, name)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	var greeting string
	if err := we.Get(context.Background(), &greeting); err != nil {
		log.Fatalln("Unable to get Workflow result", err)
	}

	fmt.Printf("WorkflowID: %s, RunID: %s, Result: %s\n", we.GetID(), we.GetRunID(), greeting)
}
