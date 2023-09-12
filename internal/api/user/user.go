package user

import (
	"fmt"
	"log"
	"mzda/internal/svc/user"
	"net/http"
)

func SignUp(svc user.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/api/user/SignUp"
		err, statusCode := svc.CreateUser(r)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, err.Error(), statusCode)
			return
		}
		w.WriteHeader(statusCode)
		return
	}
}

func ChangeUsername(svc user.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/api/user/ChangeUsername"
		err, statusCode := svc.ChangeUsername(r)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, err.Error(), statusCode)
			return
		}
		w.WriteHeader(statusCode)
		return
	}
}

func ChangePassword(svc user.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/api/user/ChangePassword"
		err, statusCode := svc.ChangePassword(r)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, err.Error(), statusCode)
			return
		}
		w.WriteHeader(statusCode)
		return
	}
}

func ChangeEmail(svc user.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/api/user/ChangePassword"
		err, statusCode := svc.ChangeEmail(r)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, err.Error(), statusCode)
			return
		}
		w.WriteHeader(statusCode)
		return
	}
}
