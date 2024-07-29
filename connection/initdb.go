package connection

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB, err = Connect()

func InitDb() {
	if err != nil {
		fmt.Println("Database connection error")
		log.Fatal(err)
	}

	if err := migrate(DB); err != nil {
		fmt.Println("Database migration error")
		log.Fatal(err)
	}

	fmt.Println("Database connected successfully")
}

func migrate(db *sql.DB) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS categories (
			id INTEGER PRIMARY KEY,
			category TEXT NOT NULL,
			category_url TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS commentdislikes (
			id INTEGER PRIMARY KEY,
			user_id INTEGER NOT NULL,
			comment_id INTEGER NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS commentlikes (
			id INTEGER PRIMARY KEY,
			user_id INTEGER NOT NULL,
			comment_id INTEGER NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS comments (
			id INTEGER PRIMARY KEY,
			post_id INTEGER NOT NULL,
			author TEXT NOT NULL,
			content TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			likes INTEGER DEFAULT 0,
			dislikes INTEGER DEFAULT 0
		);`,
		`CREATE TABLE IF NOT EXISTS dislikes (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			post_id INTEGER NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS entries (
			ID INTEGER PRIMARY KEY,
			content TEXT NOT NULL,
			categories TEXT NOT NULL,
			img_url TEXT,
			title VARCHAR(100),
			author_name TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			entry_url TEXT,
			likes INTEGER DEFAULT 0,
			dislikes INTEGER DEFAULT 0
		);`,
		`CREATE TABLE IF NOT EXISTS entrycategory (
			ID INTEGER PRIMARY KEY,
			category VARCHAR(100) NOT NULL,
			category_url VARCHAR(100) NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS likes (
			id INTEGER PRIMARY KEY,
			user_id INTEGER NOT NULL,
			post_id INTEGER NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS posts (
			id INTEGER PRIMARY KEY,
			content TEXT NOT NULL,
			categories TEXT NOT NULL,
			img_url TEXT,
			title TEXT NOT NULL,
			author_name TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			blog_url TEXT,
			likes INTEGER DEFAULT 0,
			dislikes INTEGER DEFAULT 0
		);`,
		`CREATE TABLE IF NOT EXISTS users (
			ID INTEGER PRIMARY KEY AUTOINCREMENT,
			username VARCHAR(100) NOT NULL UNIQUE,
			email VARCHAR(120) NOT NULL UNIQUE,
			password TEXT NOT NULL,
			isadmin INTEGER DEFAULT 0,
			avatar_url TEXT,
			profile_url TEXT,
			content VARCHAR(200)
		);`,
	}

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}

	return nil
}
