package auth

import (
	"fmt"
	"log"
	"mzda/internal/svc/auth"
	"net/http"
)

// SignIn
//
//	POST /auth/signin
func SignIn(svc auth.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/api/auth/SignIn"

		response, statusCode, err := svc.LoginUser(r)
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

// RenewToken
//
//	POST /auth/renew
func RenewToken(svc auth.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/api/auth/RenewToken"

		response, statusCode, err := svc.Renew(r)
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
