package user

import (
	"mzda/internal/storage/models"
	"net/http"
)

type UserSvc struct {
	userStorage models.UserStorage
}

func NewUserSvc(usr models.UserStorage) *UserSvc {
	const fn = "internal/svc/user/userSvc/NewUserSvc"
	return &UserSvc{
		userStorage: usr,
	}
}

type UserService interface {
	CreateUser(req *http.Request) (statusCode int, err error)
	ChangeUsername(req *http.Request) (statusCode int, err error)
	ChangePassword(req *http.Request) (statusCode int, err error)
	ChangeEmail(req *http.Request) (statusCode int, err error)
}
