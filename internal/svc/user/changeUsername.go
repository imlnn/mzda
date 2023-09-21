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

type ChangeUsernameRequest struct {
	Username    string `json:"username"`
	NewUsername string `json:"newUsername"`
}

func parseChangeUsername(b io.ReadCloser) (*ChangeUsernameRequest, error) {
	const fn = "internal/svc/user/changeUsername/parseChangeUsername"
	var req ChangeUsernameRequest

	err := json.NewDecoder(b).Decode(&req)
	if err != nil {
		return nil, fmt.Errorf("%s %v", fn, err)
	}

	return &req, nil
}

func (svc *Svc) ChangeUsername(req *http.Request) (statusCode int, err error) {
	const fn = "internal/svc/user/changeUsername/ChangeUsername"
	request, err := parseChangeUsername(req.Body)
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

	usr.Username = request.NewUsername
	err = svc.userStorage.UpdateUser(usr)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return http.StatusInternalServerError, fmt.Errorf("failed to update")
	}

	return http.StatusOK, nil
}
