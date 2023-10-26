package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

type Connection struct {
	db *sql.DB
}

func New() (*Connection, error) {
	const fn = "internal/storage/db/postgres/connection/New"
	var db Connection

	dbUser := os.Getenv("POSTGRES_USER")
	dbPwd := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPwd, dbName)

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
