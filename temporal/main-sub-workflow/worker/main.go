package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	app "github.com/briannqc/learning-hub/temporal/main-sub-workflow"
)

func main() {
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("Unable to connect to temporal server", err)
	}
	defer c.Close()

	w := worker.New(c, app.OnboardTaskQueue, worker.Options{})
	w.RegisterWorkflow(app.OnboardWorkflow)
	w.RegisterWorkflow(app.OnboardWithGoroutineWorkflow)
	w.RegisterWorkflow(app.HROnboardWorkflow)
	w.RegisterWorkflow(app.ITOnboardWorkflow)
	w.RegisterActivity(app.SignEmploymentAgreement)
	w.RegisterActivity(app.SignNonDisclosureAgreement)
	w.RegisterActivity(app.ActivateAccount)
	w.RegisterActivity(app.IssueLaptop)
	w.RegisterActivity(app.IssueMonitor)

	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalln("Unable to start worker", err)
	}
}
