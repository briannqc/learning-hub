package main

import (
	"log"
	"net/http"

	"go.temporal.io/sdk/client"

	emailsub "github.com/briannqc/learning-hub/temporal/email-subscription"
)

func main() {
	temporalClient, err := client.Dial(client.Options{
		HostPort:  client.DefaultHostPort,
		Namespace: client.DefaultNamespace,
	})
	if err != nil {
		log.Fatalln("Unable to dial Temporal client", err)
	}

	handler := Handler{temporalClient: temporalClient}
	http.HandleFunc("/subscribe", handler.HandleSubscribe)
	http.HandleFunc("/unsubscribe", handler.HandleUnsubscribe)
	http.HandleFunc("/details", handler.HandleGetDetails)

	if err := http.ListenAndServe(emailsub.ClientHostPort, nil); err != nil {
		log.Fatalln("Unable to start server", err)
	}
}
