package user

import (
	"mzda/internal/storage/models"
	"net/http"
)

type UserSvc struct {
	userStorage models.UserStorage
}

func NewUserSvc(usr models.UserStorage) *UserSvc {
	return &UserSvc{
		userStorage: usr,
	}
}

type UserService interface {
	CreateUser(req *http.Request) (err error, statusCode int)
	ChangeUsername(req *http.Request) (err error, statusCode int)
	ChangePassword(req *http.Request) (err error, statusCode int)
	ChangeEmail(req *http.Request) (err error, statusCode int)
}
