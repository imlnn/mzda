package postgres

import (
	"database/sql"
	"fmt"
	"log"
)

type Connection struct {
	db *sql.DB
}

func New() (*Connection, error) {
	const fn = "internal/storage/db/postgres/connector/CreateConnection"
	var db Connection

	connStr := "postgres://postgres:password@localhost/mzda"
	var err error
	db.db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Println(fmt.Errorf("%s %v", fn, err))
		return nil, err
	}

	return &db, nil
}
