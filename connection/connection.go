package connection

import (
	"database/sql"
	"forum/helpers"
	"strings"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "connection/forum.db")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ControlUsername(db *sql.DB, username string) (bool, error) {
	username = strings.TrimSpace(username)
	query := `select username from users where username = ?`
	row := db.QueryRow(query, username)
	var name string
	err := row.Scan(&name)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func ControlEmail(db *sql.DB, email string) (bool, error) {
	email = strings.TrimSpace(email)
	query := `select email from users where email = ?`
	row := db.QueryRow(query, email)
	var mail string
	err := row.Scan(&mail)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return true, nil
}

func GetUser(db *sql.DB, username string) (helpers.User, error) {
	username = strings.TrimSpace(username)
	query := `select id, username, email, avatar_url, profile_url, coalesce(content, '') as content from users where username = ?`
	row := db.QueryRow(query, username)
	var user helpers.User
	err := row.Scan(&user.ID, &user.Email, &user.AvatarUrl, &user.ProfileUrl, &user.Content)
	if err == sql.ErrNoRows {
		return helpers.User{}, nil
	} else if err != nil {
		return helpers.User{}, err
	}
	return user, nil
}
