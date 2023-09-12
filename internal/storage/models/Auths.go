package models

import (
	"time"
)

type Auth struct {
	Username     string
	RefreshToken string
	Expires      time.Time
}

func (a *Auth) IsExpired() bool {
	return a.Expires.Before(time.Now())
}

type AuthsStorage interface {
	AddAuth(auth *Auth) error
	GetAuth(token string) (*Auth, error)
	GetAuthByUser(username string) (*Auth, error)
	DeleteAuth(token string) error
	DeleteExpired() error
}
