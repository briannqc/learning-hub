package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"

	app "github.com/briannqc/learning-hub/temporal/main-sub-workflow"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to connect to temporal server", err)
	}
	defer c.Close()

	id := uuid.NewString()
	input := app.OnboardInput{
		Name: "Brian",
		ID:   id,
		Date: time.Now(),
	}
	workflowID := fmt.Sprintf("onboard-%s", input.ID)
	options := client.StartWorkflowOptions{
		ID:        workflowID,
		TaskQueue: app.OnboardTaskQueue,
	}

	we, err := c.ExecuteWorkflow(context.Background(), options, app.OnboardWithGoroutineWorkflow, input)
	if err != nil {
		log.Fatalln("Unable to execute workflow", err)
	}

	var output app.OnboardOutput
	if err := we.Get(context.Background(), &output); err != nil {
		log.Fatalln("Unable to get Workflow result", err)
	}

	fmt.Printf("WorkflowID: %s, RunID: %s, Result: %v\n", we.GetID(), we.GetRunID(), output)
}
