package user

import (
	"mzda/internal/svc/user"
	"net/http"
)

func SignUp(svc user.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/auth/api/auth/SignUp"
		err, statusCode := svc.CreateUser(r)
		if err != nil {
			http.Error(w, err.Error(), statusCode)
		}
		w.WriteHeader(statusCode)
		return
	}
}

func ChangeUsername(svc user.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/auth/api/auth/SignUp"
		err, statusCode := svc.ChangeUsername(r)
		if err != nil {
			http.Error(w, err.Error(), statusCode)
		}
		w.WriteHeader(statusCode)
		return
	}
}

func ChangePassword(svc user.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/auth/api/auth/SignUp"
		err, statusCode := svc.ChangePassword(r)
		if err != nil {
			http.Error(w, err.Error(), statusCode)
		}
		w.WriteHeader(statusCode)
		return
	}
}

func ChangeEmail(svc user.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/auth/api/auth/SignUp"
		err, statusCode := svc.ChangeEmail(r)
		if err != nil {
			http.Error(w, err.Error(), statusCode)
		}
		w.WriteHeader(statusCode)
		return
	}
}
