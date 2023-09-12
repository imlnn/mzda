package postgres

import (
	"fmt"
	"log"
	"mzda/internal/storage/models"
	"time"
)

func (c *Connection) AddAuth(auth *models.Auth) error {
	const fn = "internal/storage/db/postgres/authStorage/AddAuth"
	stmt, err := c.db.Prepare("INSERT INTO auth (username, refresh_token, expires) VALUES ($1, $2, $3)")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	_, err = stmt.Exec(auth.Username, auth.RefreshToken, auth.Expires)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	return nil
}

func (c *Connection) GetAuth(token string) (*models.Auth, error) {
	const fn = "internal/storage/db/postgres/authStorage/GetAuth"
	stmt, err := c.db.Prepare("SELECT * FROM auth WHERE refresh_token = $1")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	var auth models.Auth

	err = stmt.QueryRow(token).Scan(&auth.Username, &auth.RefreshToken, &auth.Expires)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	if auth.Username == "" {
		return nil, fmt.Errorf("auth not found")
	}

	return &auth, nil
}

func (c *Connection) GetAuthByUser(username string) (*models.Auth, error) {
	const fn = "internal/storage/db/postgres/authStorage/GetAuthByUser"
	stmt, err := c.db.Prepare("SELECT * FROM auth WHERE username = $1")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	var auth models.Auth

	err = stmt.QueryRow(username).Scan(&auth.Username, &auth.RefreshToken, &auth.Expires)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	if auth.Username == "" {
		return nil, fmt.Errorf("auth not found")
	}

	return &auth, nil
}

func (c *Connection) DeleteAuth(token string) error {
	const fn = "internal/storage/db/postgres/authStorage/DeleteAuth"
	stmt, err := c.db.Prepare("DELETE FROM auth WHERE refresh_token = $1")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	_, err = stmt.Exec(token)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	return nil
}

func (c *Connection) DeleteExpired() error {
	const fn = "internal/storage/db/postgres/authStorage/DeleteExpired"
	stmt, err := c.db.Prepare("DELETE FROM auth WHERE expires < $1")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	_, err = stmt.Exec(time.Now())
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	return nil
}
