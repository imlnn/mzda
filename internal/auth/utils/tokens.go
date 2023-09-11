package utils

import (
	"fmt"
	"log"
	"mzda/internal/storage/models/mzda"
)

func GenerateTokens(usr *mzda.User, storage mzda.AuthsStorage) (jwt string, refresh string, err error) {
	const fn = "internal/auth/handlers/auth/SignIn"
	jwt, err = GenerateJWT(usr.Username, usr.Role)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		err = fmt.Errorf("failed to generate jwt")
		return "", "", err
	}

	auth := mzda.NewAuth(usr.Username)
	err = storage.AddAuth(auth)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		err = fmt.Errorf("failed to store refresh token")
		return "", "", err
	}
	return jwt, refresh, nil
}
