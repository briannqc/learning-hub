package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	emailsub "github.com/briannqc/learning-hub/temporal/email-subscription"
)

func main() {
	c, err := client.Dial(client.Options{
		HostPort:  client.DefaultHostPort,
		Namespace: client.DefaultNamespace,
	})
	if err != nil {
		log.Fatalln("Unable to create Temporal client", err)
	}
	defer c.Close()

	w := worker.New(c, emailsub.TaskQueueName, worker.Options{})
	w.RegisterWorkflow(emailsub.SubscriptionWorkflow)
	w.RegisterActivity(emailsub.SendEmail)

	log.Println("Worker is starting")
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
