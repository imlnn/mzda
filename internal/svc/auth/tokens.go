package auth

import (
	"fmt"
	"log"
	"mzda/internal/storage/models"
	"mzda/internal/utils"
)

func (svc *AuthSvc) generateTokens(usr *models.User) (jwt string, refresh string, err error) {
	const fn = "internal/svc/auth/tokens/generateTokens"
	jwt, err = utils.GenerateJWT(usr.Username, usr.Role)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		err = fmt.Errorf("failed to generate jwt")
		return "", "", err
	}

	auth := newAuth(usr.Username)
	err = svc.authStorage.AddAuth(auth)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		err = fmt.Errorf("failed to store auth")
		return "", "", err
	}
	refresh = auth.RefreshToken
	return jwt, refresh, nil
}
