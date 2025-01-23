package database

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DB    *sql.DB
	Mutex sync.Mutex
)

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS lockers (
			id TEXT PRIMARY KEY,
			launches INTEGER DEFAULT 0,
			infections INTEGER DEFAULT 0,
			last_infection_date DATE DEFAULT NULL
		);

		CREATE TABLE IF NOT EXISTS launches (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			locker_id TEXT,
			launch_date DATE,
			FOREIGN KEY(locker_id) REFERENCES lockers(id)
		);
	`)
	if err != nil {
		log.Fatalf("Error creating tables: %v", err)
	}

	log.Println("Database initialized.")
}
