package mock

import (
	"mzda/internal/storage/models"
)

type Connection struct {
	auths []*models.Auth
	users []*models.User

	failAuth  map[string]bool
	failUsers map[string]bool
}

func NewMockConnection() *Connection {
	conn := &Connection{}

	authMethods := 5
	usersMethods := 6

	conn.failAuth = make(map[string]bool, authMethods)
	conn.failUsers = make(map[string]bool, usersMethods)
	return conn
}

func (c *Connection) FailAuthMethod(methods ...string) {
	for _, method := range methods {
		c.failAuth[method] = true
	}
}

func (c *Connection) FixAuthMethod(methods ...string) {
	for _, method := range methods {
		c.failAuth[method] = false
	}
}

func (c *Connection) FailUsersMethod(methods ...string) {
	for _, method := range methods {
		c.failUsers[method] = true
	}
}

func (c *Connection) FixUsersMethod(methods ...string) {
	for _, method := range methods {
		c.failUsers[method] = false
	}
}
