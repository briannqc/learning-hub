package main

import (
	"context"
	"log"

	"go.temporal.io/sdk/client"

	app "github.com/briannqc/learning-hub/temporal/money-transfer"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to create Temporal client:", err)
	}
	defer c.Close()

	input := app.PaymentDetails{
		SourceAccount: "85-150",
		TargetAccount: "43-812",
		Amount:        250,
		ReferenceID:   "12345",
	}

	options := client.StartWorkflowOptions{
		ID:        "pay-invoice-701",
		TaskQueue: app.MoneyTransferTaskQueueName,
	}

	log.Printf("Starting transfer from account %s to account %s for %d\n",
		input.SourceAccount, input.TargetAccount, input.Amount)
	we, err := c.ExecuteWorkflow(context.Background(), options, app.MakeTransfer, input)
	if err != nil {
		log.Fatalln("Unable to execute workflow:", err)
	}
	log.Printf("WorkflowID: %s RunID: %s\n", we.GetID(), we.GetRunID())

	var result string
	if err := we.Get(context.Background(), &result); err != nil {
		log.Fatalln("Unable to get result from workflow:", err)
	}
	log.Printf("Result: %s\n", result)
}
