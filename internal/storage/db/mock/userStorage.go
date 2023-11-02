package mock

import (
	"fmt"
	"mzda/internal/storage/models"
)

func (c *Connection) AddUser(usr *models.UserDTO) error {
	if c.failUsers["AddUser"] {
		return fmt.Errorf("failed to establish connection")
	}

	newUser := &models.User{
		ID:       1,
		Username: usr.Username,
		Pwd:      usr.Pwd,
		Email:    usr.Email,
		Role:     0,
	}
	c.users = append(c.users, newUser)
	return nil
}

func (c *Connection) UserByName(username string) (*models.User, error) {
	if c.failUsers["UserByName"] {
		return nil, fmt.Errorf("failed to establish connection")
	}

	for _, v := range c.users {
		if v.Username == username {
			return v, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func (c *Connection) UserByEmail(email string) (*models.User, error) {
	if c.failUsers["UserByEmail"] {
		return nil, fmt.Errorf("failed to establish connection")
	}

	for _, v := range c.users {
		if v.Email == email {
			return v, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func (c *Connection) UserByID(userID int) (*models.User, error) {
	if c.failUsers["UserByID"] {
		return nil, fmt.Errorf("failed to establish connection")
	}

	for _, v := range c.users {
		if v.ID == userID {
			return v, nil
		}
	}
	return nil, fmt.Errorf("user not found")
}

func (c *Connection) UpdateUser(usr *models.User) error {
	if c.failUsers["UpdateUser"] {
		return fmt.Errorf("failed to establish connection")
	}

	for k, v := range c.users {
		if v.ID == usr.ID {
			c.users[k] = usr
			return nil
		}
	}
	return fmt.Errorf("user not found")
}

func (c *Connection) DeleteUser(usr *models.User) error {
	if c.failUsers["DeleteUser"] {
		return fmt.Errorf("failed to establish connection")
	}

	for k, v := range c.users {
		if v.ID == usr.ID {
			c.users = append(c.users[:k], c.users[k+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user not found")
}
