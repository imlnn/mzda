package user

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type ChangeUsernameRequest struct {
	Username    string `json:"username"`
	NewUsername string `json:"newUsername"`
}

func parseChangeUsername(b io.ReadCloser) (*ChangeUsernameRequest, error) {
	const fn = "internal/auth/utils/users/ParseChangePassword"
	var req ChangeUsernameRequest

	err := json.NewDecoder(b).Decode(&req)
	if err != nil {
		return nil, fmt.Errorf("%s %v", fn, err)
	}

	return &req, nil
}

func (svc *UserSvc) ChangeUsername(req *http.Request) (err error, statusCode int) {
	const fn = "internal/auth/api/user/ChangeUsername"
	request, err := parseChangeUsername(req.Body)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return fmt.Errorf("failed to parse request"), http.StatusBadRequest
	}

	usr, err := svc.userStorage.UserByName(request.Username)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return fmt.Errorf("user not found"), http.StatusNotFound
	}

	usr.Username = request.NewUsername
	err = svc.userStorage.UpdateUser(usr)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return fmt.Errorf("failed to update"), http.StatusInternalServerError
	}

	return nil, http.StatusOK
}
