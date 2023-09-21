package user

import (
	"mzda/internal/storage/models"
	"net/http"
)

type Svc struct {
	userStorage models.UserStorage
}

func NewUserSvc(usr models.UserStorage) *Svc {
	const fn = "internal/svc/user/userSvc/NewUserSvc"
	return &Svc{
		userStorage: usr,
	}
}

type Service interface {
	CreateUser(req *http.Request) (statusCode int, err error)
	ChangeUsername(req *http.Request) (statusCode int, err error)
	ChangePassword(req *http.Request) (statusCode int, err error)
	ChangeEmail(req *http.Request) (statusCode int, err error)
}
