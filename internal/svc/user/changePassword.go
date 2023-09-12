package user

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	utils2 "mzda/internal/utils"
	"net/http"
)

type ChangePasswordRequest struct {
	Username    string `json:"username"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

func parseChangePassword(b io.ReadCloser) (*ChangePasswordRequest, error) {
	const fn = "internal/auth/utils/users/ParseChangePassword"
	var req ChangePasswordRequest

	err := json.NewDecoder(b).Decode(&req)
	if err != nil {
		return nil, fmt.Errorf("%s %v", fn, err)
	}

	return &req, nil
}

func (svc *UserSvc) ChangePassword(req *http.Request) (err error, statusCode int) {
	const fn = "internal/auth/api/auth/SignUp"

	request, err := parseChangePassword(req.Body)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return fmt.Errorf("failed to parse request"), http.StatusBadRequest
	}

	usr, err := svc.userStorage.UserByName(request.Username)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return fmt.Errorf("user not found"), http.StatusNotFound
	}

	if !utils2.CheckPasswordsEquality(request.OldPassword, usr.Pwd) {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return fmt.Errorf("passwords don't match"), http.StatusBadRequest
	}

	if !utils2.CheckPasswordSecurity(request.NewPassword) {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return fmt.Errorf("new password don't satifies min length == 8"), http.StatusBadRequest
	}

	usr.Pwd = request.NewPassword
	err = svc.userStorage.UpdateUser(usr)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return fmt.Errorf("failed to update"), http.StatusInternalServerError
	}

	return nil, http.StatusOK
}
