package auth

import (
	"mzda/internal/storage/models"
	"mzda/internal/utils"
	"net/http"
	"time"
)

type AuthService interface {
	LoginUser(req *http.Request) (res []byte, err error, statusCode int)
	Renew(req *http.Request) (res []byte, err error, statusCode int)
}

type AuthSvc struct {
	authStorage models.AuthsStorage
	userStorage models.UserStorage
}

func NewAuthSvc(auth models.AuthsStorage, usr models.UserStorage) *AuthSvc {
	return &AuthSvc{
		authStorage: auth,
		userStorage: usr,
	}
}

func newAuth(username string) *models.Auth {
	refreshToken := utils.GenerateRefresh()
	expires := time.Now().Add(24 * 32 * time.Hour)
	return &models.Auth{
		Username:     username,
		RefreshToken: refreshToken,
		Expires:      expires,
	}
}
