package store

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func NewSqliteStorage(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", cfg.DBName)
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}
