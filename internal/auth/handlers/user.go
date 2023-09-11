package handlers

import (
	"fmt"
	"log"
	"mzda/internal/auth/utils"
	"mzda/internal/storage/models/mzda"

	"net/http"
)

type ChangeUsernameRequest struct {
	Username    string `json:"username"`
	NewUsername string `json:"newUsername"`
}

type ChangePasswordRequest struct {
	Username    string `json:"username"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

type ChangeEmailRequest struct {
	Username string `json:"username"`
	NewEmail string `json:"newEmail"`
}

func ChangePassword(storage mzda.UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/auth/handlers/user/ChangePassword"
		req, err := utils.ParseChangePassword(r.Body)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, "failed to parse request", http.StatusBadRequest)
			return
		}

		usr, err := storage.UserByName(req.Username)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		if !utils.CheckPasswordsEquality(req.OldPassword, usr.Pwd) {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, "passwords not matching", http.StatusBadRequest)
			return
		}

		if !utils.CheckPasswordSecurity(req.NewPassword) {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, "new password don't meet security requirements", http.StatusBadRequest)
			return
		}

		usr.Pwd = req.NewPassword
		err = storage.UpdateUser(usr)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, "failed to update", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	}
}

func ChangeUsername(storage mzda.UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/auth/handlers/user/ChangeUsername"
		req, err := utils.ParseChangeUsername(r.Body)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, "failed to parse request", http.StatusBadRequest)
			return
		}

		usr, err := storage.UserByName(req.Username)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		usr.Username = req.NewUsername
		err = storage.UpdateUser(usr)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, "failed to update", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	}
}

func ChangeEmail(storage mzda.UserStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const fn = "internal/auth/handlers/user/ChangeEmail"
		req, err := utils.ParseChangeEmail(r.Body)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, "failed to parse request", http.StatusBadRequest)
			return
		}

		usr, err := storage.UserByName(req.Username)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		usr.Email = req.NewEmail
		err = storage.UpdateUser(usr)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			http.Error(w, "failed to update", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		return
	}
}
