package mzda

import (
	"mzda/internal/auth/utils"
	"time"
)

type Auth struct {
	Username     string
	RefreshToken string
	Expires      time.Time
}

func NewAuth(username string) *Auth {
	refreshToken := utils.GenerateRefresh()
	expires := time.Now().Add(24 * 32 * time.Hour)
	return &Auth{
		Username:     username,
		RefreshToken: refreshToken,
		Expires:      expires,
	}
}

func (a *Auth) IsExpired() bool {
	return a.Expires.Before(time.Now())
}

type AuthsStorage interface {
	AddAuth(auth *Auth) error
	GetAuth(token string) (*Auth, error)
	DeleteAuth(token string) error
	DeleteExpired() error
}
