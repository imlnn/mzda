package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
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

	connStr := "user=postgres password=postgrespw port=55000 sslmode=disable"

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
