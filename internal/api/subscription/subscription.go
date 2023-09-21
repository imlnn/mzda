package subscription

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"mzda/internal/svc/subscription"
	"net/http"
	"strconv"
)

// NewSubscription
//
//	POST /subscription
func NewSubscription(svc subscription.Service) http.HandlerFunc {
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
func GetSubscription(svc subscription.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/api/subscription/GetSubscription"

		param := chi.URLParam(r, "id")
		if param == "" {
			err := fmt.Errorf("%s %v", fn, "id is empty")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(param)
		if err != nil {
			err := fmt.Errorf("%s %v", fn, "id is not a number")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response, statusCode, err := svc.SubscriptionByID(id)
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
func UpdateSubscription(svc subscription.Service) http.HandlerFunc {
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
func DeleteSubscription(svc subscription.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/api/subscription/DeleteSubscription"

		param := chi.URLParam(r, "id")
		if param == "" {
			err := fmt.Errorf("%s %v", fn, "id is empty")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(param)
		if err != nil {
			err := fmt.Errorf("%s %v", fn, "id is not a number")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		statusCode, err := svc.DeleteSubscription(id, r)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, err.Error(), statusCode)
			return
		}
		w.WriteHeader(statusCode)
		return
	}
}
