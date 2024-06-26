package database

import (
	"database/sql"
	"fmt"
	"os"
)

var DB *sql.DB

func Connect() {
	var err error

	hostname := os.Getenv("POSTGRES_HOSTNAME")
	database := os.Getenv("POSTGRES_DATABASE")
	username := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")

	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", hostname, username, password, database)

	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}
}

func Close() {
	DB.Close()
}
