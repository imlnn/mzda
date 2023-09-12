package auth

import (
	"fmt"
	"log"
	"mzda/internal/svc/auth"
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

func SignIn(svc auth.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/auth/api/auth/SignIn"

		response, err, statusCode := svc.LoginUser(r)
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
			http.Error(w, "failed to send tokens", http.StatusInternalServerError)
			return
		}
		return
	}
}

func RenewToken(svc auth.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/auth/api/auth/RenewToken"

		response, err, statusCode := svc.Renew(r)
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
			http.Error(w, "failed to send tokens", http.StatusInternalServerError)
			return
		}
		return
	}
}
