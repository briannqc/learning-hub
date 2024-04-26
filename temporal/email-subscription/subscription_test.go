package emailsub

import (
	"testing"
	"time"

	"go.temporal.io/sdk/testsuite"
)

func Test_CanceledSubscriptionWorkflow(t *testing.T) {
	testSuite := &testsuite.WorkflowTestSuite{}
	env := testSuite.NewTestWorkflowEnvironment()

	details := EmailDetails{
		EmailAddress:      "me@fancyemail.com",
		Message:           "This is a test to see if the workflow cancels. This is dependent on the bool variable in the testDetails struct",
		IsSubscribed:      true,
		SubscriptionCount: 12,
	}
	env.RegisterDelayedCallback(func() {
		env.CancelWorkflow()
	}, 5*time.Second)

	env.RegisterWorkflow(SubscriptionWorkflow)
	env.RegisterActivity(SendEmail)

	env.ExecuteWorkflow(SubscriptionWorkflow, details)
}
