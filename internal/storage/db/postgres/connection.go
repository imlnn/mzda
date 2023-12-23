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

func (c *Connection) InitDB() error {
	const fn = "InitDB"
	queries := []string{
		// Subscriptions table
		`CREATE TABLE IF NOT EXISTS subscriptions(
            id serial PRIMARY KEY,
            name varchar(255) NOT NULL UNIQUE,
            description varchar(255) NOT NULL,
            admin_id int NOT NULL,
            max_members int NOT NULL CHECK(max_members > 0),
            price int NOT NULL CHECK(price > -1),
            currency varchar(3) NOT NULL,
            commission int NOT NULL,
            charge_period int NOT NULL,
            creation timestamp NOT NULL,
            start timestamp NOT NULL,
            ending timestamp
        );`,

		// Users table
		`CREATE TABLE IF NOT EXISTS users(
            id serial PRIMARY KEY,
            Username varchar(255) NOT NULL UNIQUE,
            password varchar(255) NOT NULL,
            email varchar(255) NOT NULL UNIQUE,
            role int NOT NULL
        );`,

		// Subscribers table
		`CREATE TABLE IF NOT EXISTS subscribers(
            id serial PRIMARY KEY,
            userID int NOT NULL,
            subscriptionID int NOT NULL,
            subscription_start timestamp NOT NULL,
            subscription_ending timestamp,
            FOREIGN KEY (userID) REFERENCES users (id),
            FOREIGN KEY (subscriptionID) REFERENCES subscriptions (id)
        );`,

		// Invoices table
		`CREATE TABLE IF NOT EXISTS invoices(
            id serial PRIMARY KEY,
            userID int NOT NULL,
            subscriptionID int NOT NULL,
            amount int NOT NULL,
            issued timestamp NOT NULL,
            payed timestamp,
            FOREIGN KEY (userID) REFERENCES users (id),
            FOREIGN KEY (subscriptionID) REFERENCES subscriptions (id)
        );`,

		// Payments table
		`CREATE TABLE IF NOT EXISTS payments(
            invoiceID serial PRIMARY KEY,
            amount int NOT NULL,
            accepted boolean,
            FOREIGN KEY (invoiceID) REFERENCES invoices (id)
        );`,

		// Auth table
		`CREATE TABLE IF NOT EXISTS auth(
            Username varchar(255) PRIMARY KEY,
            refresh_token varchar(10) NOT NULL,
            expires timestamp NOT NULL
        );`,
	}

	err := fmt.Errorf("")
	err = nil
	for _, query := range queries {
		_, err = c.db.Exec(query)
		if err != nil {
			log.Printf("%s: error executing query: %v", fn, err)
			return err
		}
	}

	log.Println("Database initialized successfully")
	return err
}

func (c *Connection) CleanDB() error {
	const fn = "CleanDB"
	queries := []string{
		// Dropping tables in reverse order to handle foreign key dependencies
		`DROP TABLE IF EXISTS payments CASCADE;`,
		`DROP TABLE IF EXISTS invoices CASCADE;`,
		`DROP TABLE IF EXISTS subscribers CASCADE;`,
		`DROP TABLE IF EXISTS users CASCADE;`,
		`DROP TABLE IF EXISTS subscriptions CASCADE;`,
		`DROP TABLE IF EXISTS auth CASCADE;`,
	}

	err := fmt.Errorf("")
	err = nil
	for _, query := range queries {
		_, err = c.db.Exec(query)
		if err != nil {
			log.Printf("%s: error executing query: %v", fn, err)
		}
	}

	log.Println("Database cleaned successfully")
	return err
}
