package main

import (
	"fmt"
	"goth/internal/store"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	statements := []string{
		`CREATE TABLE party (
			id INTEGER PRIMARY KEY,
			name TEXT NOT NULL,
			website TEXT NOT NULL,
			"logoPath" TEXT
		)`,
		`CREATE TABLE IF NOT EXISTS query (
			id INTEGER PRIMARY KEY,
			text TEXT NOT NULL UNIQUE
		)`,
		`CREATE TABLE "stance" (
			id INTEGER PRIMARY KEY,
			query_id INTEGER NOT NULL,
			party_id INTEGER NOT NULL,
			text TEXT,
			FOREIGN KEY (query_id) REFERENCES query (id) ON DELETE CASCADE,
    		FOREIGN KEY (party_id) REFERENCES party (id) ON DELETE CASCADE
		)`,
	}
	cfg := store.Config{
		DBName: store.Envs.DBName,
	}

	db, err := store.NewSqliteStorage(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for _, statement := range statements {
		_, err = db.Exec(statement)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("Table created")
	}

}
