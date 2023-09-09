package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"mzda/internal/storage/models/mzda"
)

type Connection struct {
	db *sql.DB
}

func New() (*Connection, error) {
	const fn = "internal/storage/db/postgres/storage/new"
	var db Connection

	//dbUsername := os.Getenv("DB_USERNAME")
	//dbPwd := os.Getenv("DB_PWD")
	//connStr := fmt.Sprintf("postgres://%s:%s@localhost/mzda", dbUsername, dbPwd)

	//connStr := "postgres://postgres:password@localhost/public?sslmode=disable"

	connStr := "user=postgres password=password port=32768 sslmode=disable"

	var err error
	db.db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	err = db.db.Ping()
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	return &db, nil
}

func (c *Connection) AddUser(usr *mzda.UserDTO) error {
	const fn = "internal/storage/db/postgres/storage/AddUser"
	stmt, err := c.db.Prepare("INSERT INTO users (username, pwd, email, role) VALUES ($1, $2, $3, $4)")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	_, err = stmt.Exec(usr.Username, usr.Pwd, usr.Email, mzda.USER)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return err
	}

	return nil
}

func (c *Connection) UserByName(username string) (usr *mzda.User, err error) {
	const fn = "internal/storage/db/postgres/storage/UserByName"
	stmt, err := c.db.Prepare("SELECT * FROM users WHERE username = $1")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	var res mzda.User

	err = stmt.QueryRow(username).Scan(&res.ID, &res.Username, &res.Pwd, &res.Email, &res.Role)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	usr = &res
	if usr == nil {
		return nil, fmt.Errorf("User not found")
	}

	return usr, nil
}

func (c *Connection) UserByEmail(email string) (usr *mzda.User, err error) {
	const fn = "internal/storage/db/postgres/storage/UserByEmail"
	stmt, err := c.db.Prepare("SELECT * FROM users WHERE email = $1")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	var res mzda.User

	err = stmt.QueryRow(email).Scan(&res.ID, &res.Username, &res.Pwd, &res.Email, &res.Role)

	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	usr = &res
	if usr == nil {
		return nil, fmt.Errorf("User not found")
	}

	return usr, nil
}

func (c *Connection) UserByID(userID int) (usr *mzda.User, err error) {
	const fn = "internal/storage/db/postgres/storage/UserByEmail"
	stmt, err := c.db.Prepare("SELECT * FROM users WHERE id = $1")
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	var res mzda.User

	err = stmt.QueryRow(userID).Scan(&res.ID, &res.Username, &res.Pwd, &res.Email, &res.Role)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	usr = &res
	if usr == nil {
		return nil, fmt.Errorf("User not found")
	}

	return usr, nil
}

func (c *Connection) DeleteUser(usr *mzda.User) error {
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

func (c *Connection) UpdateUser(usr *mzda.User) error {
	const fn = "internal/storage/db/postgres/storage/AddUser"
	stmt, err := c.db.Prepare("UPDATE users SET username = $1, pwd = $2, email = $3 WHERE id = $4")
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
