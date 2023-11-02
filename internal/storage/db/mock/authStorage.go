package mock

import (
	"fmt"
	"mzda/internal/storage/models"
	"time"
)

func (c *Connection) AddAuth(auth *models.Auth) error {
	if c.failAuth["AddAuth"] {
		return fmt.Errorf("failed to establish connection")
	}

	c.auths = append(c.auths, auth)
	return nil
}

func (c *Connection) GetAuth(token string) (*models.Auth, error) {
	if c.failAuth["GetAuth"] {
		return nil, fmt.Errorf("failed to establish connection")
	}

	for _, v := range c.auths {
		if v.RefreshToken == token {
			return v, nil
		}
	}
	return nil, fmt.Errorf("auth not found")
}

func (c *Connection) GetAuthByUser(username string) (*models.Auth, error) {
	if c.failAuth["GetAuthByUser"] {
		return nil, fmt.Errorf("failed to establish connection")
	}

	for _, v := range c.auths {
		if v.Username == username {
			return v, nil
		}
	}
	return nil, fmt.Errorf("auth not found")
}

func (c *Connection) DeleteAuth(token string) error {
	if c.failAuth["DeleteAuth"] {
		return fmt.Errorf("failed to establish connection")
	}

	for k, v := range c.auths {
		if v.RefreshToken == token {
			c.auths = append(c.auths[:k], c.auths[k+1:]...)
			return nil
		}
	}
	return fmt.Errorf("auth not found")
}

func (c *Connection) DeleteExpired() error {
	if c.failAuth["DeleteExpired"] {
		return fmt.Errorf("failed to establish connection")
	}

	for k, v := range c.auths {
		if v.Expires.Before(time.Now()) {
			c.auths = append(c.auths[:k], c.auths[k+1:]...)
		}
	}
	return nil
}
