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

func (svc *AuthSvc) Renew(req *http.Request) (res []byte, statusCode int, err error) {
	const fn = "internal/svc/auth/renew/Renew"

	refresh := req.Header.Get("refreshToken")
	if strings.EqualFold(refresh, "") {
		err := fmt.Errorf("missing refresh token")
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, http.StatusBadRequest, err
	}

	auth, err := svc.authStorage.GetAuth(refresh)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, http.StatusNotFound, fmt.Errorf("failed to find session")
	}

	if auth.IsExpired() {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, http.StatusUnauthorized, fmt.Errorf("token is expired")
	}

	usr, err := svc.userStorage.UserByName(auth.Username)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, http.StatusNotFound, fmt.Errorf("user not found")
	}

	err = svc.authStorage.DeleteAuth(auth.RefreshToken)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, http.StatusInternalServerError, fmt.Errorf("failed to store auth")
	}

	jwt, refresh, err := svc.generateTokens(usr)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, http.StatusInternalServerError, err
	}

	response := renewResponse{
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
