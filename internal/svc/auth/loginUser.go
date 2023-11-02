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

	if req.Username == "" || req.Password == "" {
		return nil, fmt.Errorf("%s  provided JSON is empty", fn)
	}

	return &req, nil
}

func (svc *Svc) LoginUser(req *http.Request) (res []byte, statusCode int, err error) {
	const fn = "internal/svc/auth/loginUser/LoginUser"

	cred, err := parseCredentials(req.Body)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, http.StatusBadRequest, fmt.Errorf("failed to parse request")
	}

	usr, err := svc.userStorage.UserByName(cred.Username)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, http.StatusBadRequest, fmt.Errorf("user not found")
	}

	if !utils.CheckPasswordsEquality(usr.Pwd, cred.Password) {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, http.StatusUnauthorized, fmt.Errorf("passwords not match")
	}

	auth, err := svc.authStorage.GetAuthByUser(usr.Username)
	if auth != nil {
		err = svc.authStorage.DeleteAuth(auth.RefreshToken)
		if err != nil {
			log.Println(fmt.Errorf("%s %v", fn, err))
			return nil, http.StatusInternalServerError, fmt.Errorf("failed to store auth")
		}
	}

	jwt, refresh, err := svc.generateTokens(usr)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, http.StatusInternalServerError, err
	}

	response := loginResponse{
		JWT:     jwt,
		Refresh: refresh,
	}

	payload, err := json.Marshal(&response)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to generate response")
	}

	return payload, http.StatusOK, nil
}
