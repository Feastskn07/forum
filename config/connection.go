package connection

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // SQLite3 driverını içe aktar
)

func ConnectionDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./database/forum.db")
	if err != nil {
		return nil, err 
	}
	return db, nil 
}
