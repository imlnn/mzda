package auth

import (
	"mzda/internal/storage/models"
	"mzda/internal/utils"
	"net/http"
	"time"
)

type AuthService interface {
	LoginUser(req *http.Request) (res []byte, statusCode int, err error)
	Renew(req *http.Request) (res []byte, statusCode int, err error)
}

type AuthSvc struct {
	authStorage models.AuthsStorage
	userStorage models.UserStorage
}

func NewAuthSvc(auth models.AuthsStorage, usr models.UserStorage) *AuthSvc {
	const fn = "internal/svc/auth/authSvc/NewAuthSvc"
	return &AuthSvc{
		authStorage: auth,
		userStorage: usr,
	}
}

func newAuth(username string) *models.Auth {
	const fn = "internal/svc/auth/authSvc/newAuth"
	refreshToken := utils.GenerateRefresh()
	expires := time.Now().Add(24 * 32 * time.Hour)
	return &models.Auth{
		Username:     username,
		RefreshToken: refreshToken,
		Expires:      expires,
	}
}
