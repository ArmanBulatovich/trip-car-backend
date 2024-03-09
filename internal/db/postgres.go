package db

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB

func ConnectToPostgres(url string) error {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return fmt.Errorf("postgres sql.Open %s: %s", url, err.Error())
	}

	DB = db

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("postgres Ping: %s", err.Error())
	}

	return nil
}

func Close() {
	err := DB.Close()
	if err != nil {
		fmt.Printf("postgres: Close: %s\n", err.Error())
	}
}
