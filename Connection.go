// vim: sw=4 ts=4 expandtab
package main

import (
	"database/sql"
	"fmt"
	"os"

    _ "github.com/go-sql-driver/mysql"
)

func connection() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbinfo := fmt.Sprintf("%s:%s@tcp(localhost:3306)/%s", dbUser, dbPassword, dbName)

	db, db_err := sql.Open("mysql", dbinfo)
	if db_err != nil {
		return nil, db_err
	}
	return db, nil
}
