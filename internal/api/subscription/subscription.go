package subscription

import (
	"fmt"
	"log"
	"mzda/internal/svc/subscription"
	"net/http"
)

// NewSubscription
//
//	POST /subscription
func NewSubscription(svc subscription.SubscriptionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/api/subscription/NewSubscription"
		statusCode, err := svc.AddSubscription(r)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, err.Error(), statusCode)
			return
		}
		w.WriteHeader(statusCode)
		return
	}
}

// GetSubscription
//
//	GET /subscription/{id}
func GetSubscription(svc subscription.SubscriptionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/api/subscription/GetSubscription"
		response, statusCode, err := svc.SubscriptionByID(r)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, err.Error(), statusCode)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(response)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, "failed to send data", http.StatusInternalServerError)
			return
		}
		return
	}
}

// UpdateSubscription
//
//	PUT /subscription/{id}
func UpdateSubscription(svc subscription.SubscriptionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/api/subscription/UpdateSubscription"
		statusCode, err := svc.UpdateSubscription(r)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, err.Error(), statusCode)
			return
		}
		w.WriteHeader(statusCode)
		return
	}
}

// DeleteSubscription
//
//	DELETE /subscription/{id}
func DeleteSubscription(svc subscription.SubscriptionService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/api/subscription/DeleteSubscription"
		statusCode, err := svc.DeleteSubscription(r)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, err.Error(), statusCode)
			return
		}
		w.WriteHeader(statusCode)
		return
	}
}
