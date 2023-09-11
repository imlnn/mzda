package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"mzda/internal/auth/utils"
	"mzda/internal/storage/models/mzda"
	"net/http"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	JWT     string `json:"jwt"`
	Refresh string `json:"refresh"`
}

func NewAuthResponse(jwt string, refresh string) *AuthResponse {
	return &AuthResponse{
		JWT:     jwt,
		Refresh: refresh,
	}
}

func SignUp(storage mzda.UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/auth/handlers/auth/SignUp"
		usr, err := utils.ParseUserDTO(r.Body)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, "failed to parse request", http.StatusBadRequest)
			return
		}
		err = storage.AddUser(usr)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, "failed to register", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}
}

func SignIn(userStorage mzda.UserStorage, authsStorage mzda.AuthsStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/auth/handlers/auth/SignIn"

		cred, err := utils.ParseCredentials(r.Body)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, "failed to parse request", http.StatusBadRequest)
			return
		}

		usr, err := userStorage.UserByName(cred.Username)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, "failed to find user", http.StatusNotFound)
			return
		}

		if !utils.CheckPasswordsEquality(usr.Pwd, cred.Password) {
			log.Println(fmt.Errorf("%s %v", fn, fmt.Errorf("passwords not match")))
			http.Error(w, "passwords not match", http.StatusUnauthorized)
			return
		}

		jwt, refresh, err := utils.GenerateTokens(usr, authsStorage)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		payload := NewAuthResponse(jwt, refresh)
		response, err := json.Marshal(&payload)
		if err != nil {
			fmt.Println(err)
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

func RenewToken(userStorage mzda.UserStorage, authsStorage mzda.AuthsStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/auth/handlers/auth/RenewToken"

		refresh := r.Header.Get("refreshToken")
		auth, err := authsStorage.GetAuth(refresh)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, "failed to find session", http.StatusNotFound)
			return
		}

		if auth.IsExpired() {
			log.Println(fmt.Errorf("%s %v", fn, "token is expired"))
			http.Error(w, "token is expired", http.StatusUnauthorized)
			return
		}

		usr, err := userStorage.UserByName(auth.Username)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, "failed to find user", http.StatusNotFound)
			return
		}

		jwt, refresh, err := utils.GenerateTokens(usr, authsStorage)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		payload := NewAuthResponse(jwt, refresh)
		response, err := json.Marshal(&payload)
		if err != nil {
			fmt.Println(err)
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
