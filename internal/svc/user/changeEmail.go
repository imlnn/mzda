package user

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ChangeEmailRequest struct {
	Username string `json:"username"`
	NewEmail string `json:"newEmail"`
}

func parseChangeEmail(b io.ReadCloser) (*ChangeEmailRequest, error) {
	const fn = "internal/auth/utils/users/ParseChangePassword"
	var req ChangeEmailRequest

	err := json.NewDecoder(b).Decode(&req)
	if err != nil {
		return nil, fmt.Errorf("%s %v", fn, err)
	}

	return &req, nil
}

func (svc *UserSvc) ChangeEmail(req *http.Request) (err error, statusCode int) {
	const fn = "internal/auth/api/user/ChangeEmail"
	request, err := parseChangeEmail(req.Body)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return fmt.Errorf("failed to parse request"), http.StatusBadRequest
	}

	usr, err := svc.userStorage.UserByName(request.Username)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return fmt.Errorf("user not found"), http.StatusNotFound
	}

	usr.Email = request.NewEmail
	err = svc.userStorage.UpdateUser(usr)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return fmt.Errorf("failed to update"), http.StatusInternalServerError
	}

	return nil, http.StatusOK
}
