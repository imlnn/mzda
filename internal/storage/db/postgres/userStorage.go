package postgres

import (
	"fmt"
	"log"
	"mzda/internal/storage/models"
)

func (c *Connection) AddUser(usr *models.UserDTO) error {
	const fn = "internal/storage/db/postgres/storage/AddUser"
	stmt, err := c.db.Prepare("INSERT INTO users (username, password, email, role) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	_, err = stmt.Exec(usr.Username, usr.Pwd, usr.Email, models.USER)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	return nil
}

func (c *Connection) UserByName(username string) (*models.User, error) {
	const fn = "internal/storage/db/postgres/storage/UserByName"
	stmt, err := c.db.Prepare("SELECT * FROM users WHERE username = $1")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	var usr models.User

	err = stmt.QueryRow(username).Scan(&usr.ID, &usr.Username, &usr.Pwd, &usr.Email, &usr.Role)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	if usr.Username == "" {
		return nil, fmt.Errorf("User not found")
	}

	return &usr, nil
}

func (c *Connection) UserByEmail(email string) (*models.User, error) {
	const fn = "internal/storage/db/postgres/storage/UserByEmail"
	stmt, err := c.db.Prepare("SELECT * FROM users WHERE email = $1")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	var usr models.User

	err = stmt.QueryRow(email).Scan(&usr.ID, &usr.Username, &usr.Pwd, &usr.Email, &usr.Role)

	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	if usr.Username == "" {
		return nil, fmt.Errorf("User not found")
	}

	return &usr, nil
}

func (c *Connection) UserByID(userID int) (*models.User, error) {
	const fn = "internal/storage/db/postgres/storage/UserByEmail"
	stmt, err := c.db.Prepare("SELECT * FROM users WHERE id = $1")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	var usr models.User

	err = stmt.QueryRow(userID).Scan(&usr.ID, &usr.Username, &usr.Pwd, &usr.Email, &usr.Role)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	if usr.Username == "" {
		return nil, fmt.Errorf("User not found")
	}

	return &usr, nil
}

func (c *Connection) DeleteUser(usr *models.User) error {
	const fn = "internal/storage/db/postgres/storage/UserByEmail"
	stmt, err := c.db.Prepare("DELETE FROM users WHERE id = $1")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	_, err = stmt.Exec(usr.ID)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	return nil
}

func (c *Connection) UpdateUser(usr *models.User) error {
	const fn = "internal/storage/db/postgres/storage/AddUser"
	stmt, err := c.db.Prepare("UPDATE users SET username = $1, password = $2, email = $3 WHERE id = $4")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	_, err = stmt.Exec(usr.Username, usr.Pwd, usr.Email, usr.ID)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	return nil
}
