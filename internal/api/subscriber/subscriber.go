package subscriber

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"mzda/internal/svc/subscriber"
	"net/http"
	"strconv"
)

func NewSubscriber(svc subscriber.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/api/subscriber/NewSubscriber"
		statusCode, err := svc.AddSubscriber(r)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, err.Error(), statusCode)
			return
		}
		w.WriteHeader(statusCode)
		return
	}
}

func GetSubscriber(svc subscriber.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/api/subscriber/GetSubscriber"

		param := chi.URLParam(r, "id")
		if param == "" {
			err := fmt.Errorf("%s %v", fn, "id is empty")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(param)
		if err != nil {
			e := fmt.Errorf("%s %v", fn, "id is not a number")
			http.Error(w, e.Error(), http.StatusBadRequest)
			return
		}

		response, statusCode, err := svc.SubscriberByID(id)
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

func GetSubscribersListByUserID(svc subscriber.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/api/subscriber/GetSubscribers"

		param := chi.URLParam(r, "id")
		if param == "" {
			err := fmt.Errorf("%s %v", fn, "id is empty")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(param)
		if err != nil {
			e := fmt.Errorf("%s %v", fn, "id is not a number")
			http.Error(w, e.Error(), http.StatusBadRequest)
			return
		}

		response, statusCode, err := svc.ListSubscribersByUserID(id)
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

func UpdateSubscriber(svc subscriber.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/api/subscriber/UpdateSubscriber"
		statusCode, err := svc.UpdateSubscriber(r)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, err.Error(), statusCode)
			return
		}
		w.WriteHeader(statusCode)
		return
	}
}

func DeleteSubscriber(svc subscriber.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/api/subscriber/DeleteSubscriber"

		param := chi.URLParam(r, "id")
		if param == "" {
			err := fmt.Errorf("%s %v", fn, "id is empty")
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(param)
		if err != nil {
			e := fmt.Errorf("%s %v", fn, "id is not a number")
			http.Error(w, e.Error(), http.StatusBadRequest)
			return
		}

		statusCode, err := svc.DeleteSubscriberByID(id, r)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, err.Error(), statusCode)
			return
		}
		w.WriteHeader(statusCode)
		return
	}
}
