package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

var (
	Sessions     = map[string]session{}
	UserSessions = map[string]string{}
)

type session struct {
	Username string
	Email    string
	Expiry   time.Time
	IsAdmin  bool
}

func (s session) IsExpired() bool {
	return s.Expiry.Before(time.Now())
}

func CreateCookie(w http.ResponseWriter, username, email string, isAdmin bool) {
	if oldSessionToken, exists := UserSessions[email]; exists {
		delete(Sessions, oldSessionToken)
	}

	sessionToken := uuid.NewString()
	expiresAt := time.Now().Add(1 * time.Hour)
	Sessions[sessionToken] = session{
		Username: username,
		Email:    email,
		Expiry:   expiresAt,
		IsAdmin:  isAdmin,
	}

	UserSessions[email] = sessionToken
	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Path:     "/",
		HttpOnly: true,
		Expires:  expiresAt,
	})
}

func GetSession(r *http.Request) (*session, error) {
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return nil, fmt.Errorf("cookie not found: %v", err)
	}

	sessionToken := cookie.Value
	userSession, exists := Sessions[sessionToken]
	if !exists || userSession.IsExpired() {
		return nil, fmt.Errorf("session not found or expired")
	}
	return &userSession, nil
}

func SessionCheck(r *http.Request) bool {
	_, err := GetSession(r)
	return err == nil
}

func UpdateCookie(r *http.Request, username string) bool {
	userSession, err := GetSession(r)
	if err != nil {
		fmt.Println("Error updating cookie:", err)
		return false
	}

	userSession.Username = username
	return true
}

func CheckAdminSession(r *http.Request) bool {
	userSession, err := GetSession(r)
	if err != nil {
		return false
	}
	return userSession.IsAdmin
}

func RequireAdmin(w http.ResponseWriter, r *http.Request) bool {
	if !CheckAdminSession(r) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return false
	}
	return true
}
