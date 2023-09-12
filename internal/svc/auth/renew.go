package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type renewResponse struct {
	JWT     string `json:"jwt"`
	Refresh string `json:"refresh"`
}

func (svc *AuthSvc) Renew(req *http.Request) (res []byte, err error, statusCode int) {
	const fn = "internal/auth/api/auth/SignUp"

	refresh := req.Header.Get("refreshToken")
	if strings.EqualFold(refresh, "") {
		err := fmt.Errorf("missing refresh token")
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err, http.StatusBadRequest
	}

	auth, err := svc.authStorage.GetAuth(refresh)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, fmt.Errorf("failed to find session"), http.StatusNotFound
	}

	if auth.IsExpired() {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, fmt.Errorf("token is expired"), http.StatusUnauthorized
	}

	usr, err := svc.userStorage.UserByName(auth.Username)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, fmt.Errorf("user not found"), http.StatusNotFound
	}

	err = svc.authStorage.DeleteAuth(auth.RefreshToken)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, fmt.Errorf("failed to store auth"), http.StatusInternalServerError
	}

	jwt, refresh, err := svc.generateTokens(usr)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err, http.StatusInternalServerError
	}

	response := renewResponse{
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
