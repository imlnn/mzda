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

type ChangeEmailRequest struct {
	Username string `json:"username"`
	NewEmail string `json:"newEmail"`
}

func parseChangeEmail(b io.ReadCloser) (*ChangeEmailRequest, error) {
	const fn = "internal/svc/user/changeEmail/parseChangeEmail"
	var req ChangeEmailRequest

	err := json.NewDecoder(b).Decode(&req)
	if err != nil {
		return nil, fmt.Errorf("%s %v", fn, err)
	}

	return &req, nil
}

func (svc *UserSvc) ChangeEmail(req *http.Request) (err error, statusCode int) {
	const fn = "internal/svc/user/changeEmail/ChangeEmail"
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
	jwt := req.Context().Value("jwt").(*utils.JWT)
	if !strings.EqualFold(jwt.Username, usr.Username) {
		err = fmt.Errorf("token given for another user")
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err, http.StatusUnauthorized
	}

	usr.Email = request.NewEmail
	err = svc.userStorage.UpdateUser(usr)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return fmt.Errorf("failed to update"), http.StatusInternalServerError
	}

	return nil, http.StatusOK
}
