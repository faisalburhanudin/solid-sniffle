package database

import (
	"database/sql"
	"log"
)

// NewSqliteMemory Create new connection to sqlite in memory database
func NewSqliteMemory() *sql.DB {
	db, err := sql.Open("sqlite3", "file::memory:?mode=memory&cache=shared")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
