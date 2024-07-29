package connection

import (
	"database/sql"
	"fmt"
	"forum/helpers"
	"strings"
)

func GetProfile(db *sql.DB, user_url string) (helpers.User, error) {
	trimmedUserURL := strings.TrimSpace(user_url)
	query := `select id, username, email, avatar_url, profile_url, coalesce(content, '')
			  as content from users where profile_url = ?`
	row := db.QueryRow(query, trimmedUserURL)

	var id int
	var avatar_url, username, email, profile_url, content string
	err := row.Scan(&id, &username, &email, &avatar_url, &profile_url, &content)

	if err != nil {
		if err == sql.ErrNoRows {
			return helpers.User{}, nil
		}
		return helpers.User{}, err
	}

	return helpers.User{
		ID:         id,
		AvatarUrl:  avatar_url,
		ProfileUrl: profile_url,
		Username:   username,
		Email:      email,
		Content:    content,
	}, nil
}

func UpdateProfilePhoto(db *sql.DB, url_path, username string) error {
	trimmedURLPath := strings.TrimSpace(url_path)
	trimmedUsername := strings.TrimSpace(username)
	query := `update users set avatar_url = ? where username = ?`
	statement, err := db.Prepare(query)

	if err != nil {
		return fmt.Errorf("failed to prepare query: %v", err)
	}
	defer statement.Close()

	_, err = statement.Exec(trimmedURLPath, trimmedUsername)

	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}
	return nil
}

func UpdateUserInformation(db *sql.DB, username, content string) error {
	trimmedUsername := strings.TrimSpace(username)
	trimmedContent := strings.TrimSpace(content)
	profile_url := helpers.CreateUrl(trimmedUsername)
	query := `update users set username = ?, content = ?, profile_url = ? where username = ?`
	statement, err := db.Prepare(query)

	if err != nil {
		return fmt.Errorf("failed to prepare query: %v", err)
	}
	defer statement.Close()

	_, err = statement.Exec(trimmedUsername, trimmedContent, profile_url,
		helpers.SendData.Data["userinfo"].(helpers.User).Username)

	if err != nil {
		return fmt.Errorf("failed to execute query: %v", err)
	}
	return nil
}
