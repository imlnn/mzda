package user

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mzda/internal/utils"
	"net/http"
	"strings"
)

type ChangePasswordRequest struct {
	Username    string `json:"username"`
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

func parseChangePassword(b io.ReadCloser) (*ChangePasswordRequest, error) {
	const fn = "internal/svc/user/changePassword/parseChangePassword"
	var req ChangePasswordRequest

	err := json.NewDecoder(b).Decode(&req)
	if err != nil {
		return nil, fmt.Errorf("%s %v", fn, err)
	}

	return &req, nil
}

func (svc *Svc) ChangePassword(req *http.Request) (statusCode int, err error) {
	const fn = "internal/svc/user/changePassword/ChangePassword"

	request, err := parseChangePassword(req.Body)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return http.StatusBadRequest, fmt.Errorf("failed to parse request")
	}

	usr, err := svc.userStorage.UserByName(request.Username)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return http.StatusNotFound, fmt.Errorf("user not found")
	}
	jwt := req.Context().Value("jwt").(*utils.JWT)
	if !strings.EqualFold(jwt.Username, usr.Username) {
		err = fmt.Errorf("token given for another user")
		log.Println(fmt.Errorf("%s %v", fn, err))
		return http.StatusUnauthorized, err
	}

	if !utils.CheckPasswordsEquality(request.OldPassword, usr.Pwd) {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return http.StatusBadRequest, fmt.Errorf("passwords don't match")
	}

	if !utils.CheckPasswordSecurity(request.NewPassword) {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return http.StatusBadRequest, fmt.Errorf("new password don't satifies min length == 8")
	}

	usr.Pwd = request.NewPassword
	err = svc.userStorage.UpdateUser(usr)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return http.StatusInternalServerError, fmt.Errorf("failed to update")
	}

	return http.StatusOK, nil
}
