package db

import (
	"database/sql"
	"fmt"

	"github.com/pressly/goose/v3"
)

func Migrate(db *sql.DB) error {
	if err := goose.SetDialect("sqlite3"); err != nil {
		return err
	}

	if err := goose.Up(db, "migrations"); err != nil {
		return err
	}

	fmt.Println("Migrations applied")
	return nil
}