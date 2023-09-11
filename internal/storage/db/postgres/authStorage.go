package postgres

import (
	"fmt"
	"log"
	"mzda/internal/storage/models/mzda"
	"time"
)

func (c *Connection) AddAuth(auth *mzda.Auth) error {
	const fn = "internal/storage/db/postgres/storage/AddUser"
	stmt, err := c.db.Prepare("INSERT INTO auth (username, refresh_token, expires) VALUES ($1, $2, $3)")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	_, err = stmt.Exec(auth.Username, auth.RefreshToken, auth.Expires, mzda.USER)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	return nil
}

func (c *Connection) GetAuth(token string) (*mzda.Auth, error) {
	const fn = "internal/storage/db/postgres/storage/UserByName"
	stmt, err := c.db.Prepare("SELECT * FROM auth WHERE refresh_token = $1")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	var auth mzda.Auth

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

func (c *Connection) DeleteAuth(token string) error {
	const fn = "internal/storage/db/postgres/storage/UserByEmail"
	stmt, err := c.db.Prepare("DELETE FROM auth WHERE refresh_token < $1")
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
	const fn = "internal/storage/db/postgres/storage/UserByEmail"
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
