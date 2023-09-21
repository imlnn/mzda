package user

import (
	"fmt"
	"log"
	"mzda/internal/svc/user"
	"net/http"
)

// SignUp
//
//	POST /signup
func SignUp(svc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/api/user/SignUp"
		statusCode, err := svc.CreateUser(r)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, err.Error(), statusCode)
			return
		}
		w.WriteHeader(statusCode)
		return
	}
}

// ChangeUsername
//
//	POST /user/changeUsername
func ChangeUsername(svc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/api/user/ChangeUsername"
		statusCode, err := svc.ChangeUsername(r)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, err.Error(), statusCode)
			return
		}
		w.WriteHeader(statusCode)
		return
	}
}

// ChangePassword
//
//	POST /user/changePassword
func ChangePassword(svc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/api/user/ChangePassword"
		statusCode, err := svc.ChangePassword(r)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, err.Error(), statusCode)
			return
		}
		w.WriteHeader(statusCode)
		return
	}
}

// ChangeEmail
//
//	POST /user/changeEmail
func ChangeEmail(svc user.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/api/user/ChangePassword"
		statusCode, err := svc.ChangeEmail(r)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, err.Error(), statusCode)
			return
		}
		w.WriteHeader(statusCode)
		return
	}
}
