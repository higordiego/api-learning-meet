package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// singleton

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/meetings")

	if err != nil {
		return nil, err
	}

	fmt.Println()

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
