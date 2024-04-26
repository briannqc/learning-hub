package emailsub

import (
	"context"

	"go.temporal.io/sdk/activity"
)

func SendEmail(ctx context.Context, details EmailDetails) (string, error) {
	activity.GetLogger(ctx).Info("Sending email to customer", "EmailAddress", details.EmailAddress)
	return "Email sent to " + details.EmailAddress, nil
}
