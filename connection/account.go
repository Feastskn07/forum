package connection

import (
	"database/sql"
	"fmt"
	"forum/helpers"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func SetAccount(db *sql.DB, username, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	rand.Seed(time.Now().UnixNano())
	avatarURL := fmt.Sprintf("../avatars/avatar%d.png", rand.Intn(10)+1)
	profileURL := helpers.CreateUrl(username)

	query := `insert into users (username, email, password, avatar_url, profile_url) 
			  values (?, ?, ?, ?, ?)`

	_, err = db.Exec(query, strings.TrimSpace(username), strings.TrimSpace(email),
		hashedPassword, avatarURL, profileURL)

	if err != nil {
		return err
	}
	return nil
}

func GetAccount(db *sql.DB, email, password string) (string, map[string]string, error) {
	var hashedPassword, username string
	errorMessages := make(map[string]string)
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@(gmail\.com|outlook\.com)$`)

}
