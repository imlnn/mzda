package auth

import (
	"mzda/internal/storage/models/auth"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) (*auth.User, error) {

	return nil, nil
}

func ChangePassword(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ChangeUsername(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ChangeEmail(w http.ResponseWriter, r *http.Request) error {
	return nil
}
