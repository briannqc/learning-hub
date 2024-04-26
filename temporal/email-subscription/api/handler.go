package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"go.temporal.io/sdk/client"

	emailsub "github.com/briannqc/learning-hub/temporal/email-subscription"
)

type Request struct {
	Email string `json:"email"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Handler struct {
	temporalClient client.Client
}

func (h *Handler) HandleSubscribe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Invalid Content-Type, expecting application/json", http.StatusUnsupportedMediaType)
		return
	}

	var request Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if request.Email == "" {
		http.Error(w, "Missing email", http.StatusBadRequest)
		return
	}

	workflowOptions := client.StartWorkflowOptions{
		ID:                                       request.Email,
		TaskQueue:                                emailsub.TaskQueueName,
		WorkflowExecutionErrorWhenAlreadyStarted: true,
	}

	subscription := emailsub.EmailDetails{
		EmailAddress:      request.Email,
		Message:           "Welcome to the Subscription Workflow",
		IsSubscribed:      true,
		SubscriptionCount: 0,
	}

	_, err := h.temporalClient.ExecuteWorkflow(context.Background(), workflowOptions, emailsub.SubscriptionWorkflow, subscription)
	if err != nil {
		http.Error(w, "Could NOT sign up user. Please try again.", http.StatusInternalServerError)
		log.Printf("failed to execute workflow: %v", err)
		return
	}

	response := Response{
		Status:  "Success",
		Message: "Signed up",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Print("Could not encode response JSON", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (h *Handler) HandleUnsubscribe(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Invalid Content-Type, expecting application/json", http.StatusUnsupportedMediaType)
		return
	}

	var request Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if request.Email == "" {
		http.Error(w, "Missing email", http.StatusBadRequest)
		return
	}

	workflowID := request.Email
	if err := h.temporalClient.CancelWorkflow(context.Background(), workflowID, ""); err != nil {
		http.Error(w, "Couldn't unsubscribe. Please try again.", http.StatusInternalServerError)
		return
	}

	response := Response{
		Status:  "Success",
		Message: "Unsubscribed",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Print("Could not encode response JSON", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (h *Handler) HandleGetDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	queryValues, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		http.Error(w, "Couldn't query values. Please try again.", http.StatusInternalServerError)
		log.Println("Failed to query Workflow.")
		return
	}
	email := queryValues.Get("email")
	workflowID := email
	queryType := "GetDetails"

	resp, err := h.temporalClient.QueryWorkflow(context.Background(), workflowID, "", queryType)
	if err != nil {
		http.Error(w, "Couldn't query values. Please try again.", http.StatusInternalServerError)
		log.Println("Failed to query Workflow.", err)
		return
	}

	var result emailsub.EmailDetails
	if err := resp.Get(&result); err != nil {
		http.Error(w, "Couldn't query values. Please try again.", http.StatusInternalServerError)
		log.Println("Failed to query Workflow.", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Print("Could not encode response JSON", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
