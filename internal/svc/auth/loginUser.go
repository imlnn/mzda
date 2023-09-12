package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mzda/internal/utils"
	"net/http"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	JWT     string `json:"jwt"`
	Refresh string `json:"refresh"`
}

func parseCredentials(b io.ReadCloser) (*loginRequest, error) {
	const fn = "internal/svc/auth/loginUser/parseCredentials"
	var req loginRequest

	err := json.NewDecoder(b).Decode(&req)
	if err != nil {
		return nil, fmt.Errorf("%s  %v", fn, err)
	}

	return &req, nil
}

func (svc *AuthSvc) LoginUser(req *http.Request) (res []byte, err error, statusCode int) {
	const fn = "internal/svc/auth/loginUser/LoginUser"

	cred, err := parseCredentials(req.Body)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, fmt.Errorf("failed to parse request"), http.StatusBadRequest
	}

	usr, err := svc.userStorage.UserByName(cred.Username)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, fmt.Errorf("user not found"), http.StatusBadRequest
	}

	if !utils.CheckPasswordsEquality(usr.Pwd, cred.Password) {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, fmt.Errorf("passwords not match"), http.StatusUnauthorized
	}

	auth, err := svc.authStorage.GetAuthByUser(usr.Username)
	if auth != nil {
		err = svc.authStorage.DeleteAuth(auth.RefreshToken)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			return nil, fmt.Errorf("failed to store auth"), http.StatusInternalServerError
		}
	}

	jwt, refresh, err := svc.generateTokens(usr)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err, http.StatusInternalServerError
	}

	response := loginResponse{
		JWT:     jwt,
		Refresh: refresh,
	}

	payload, err := json.Marshal(&response)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, fmt.Errorf("failed to generate response"), http.StatusInternalServerError
	}

	return payload, nil, http.StatusOK
}
