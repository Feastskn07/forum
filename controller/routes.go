package routes

import (
	"fmt"
	"forum/controller"
	"forum"
	oauth "forum/controller/oAuth"
	"net/http"
)

var SendData = helpres.SendData{
	Data: make(map[string]interface{}), // Haritayı başlatıyoruz
}

// HandleRequests handles all incoming HTTP requests and routes them to the appropriate controller functions.
func HandleRequests() {
	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/logout", handlers.Logout)
	http.HandleFunc("/register", handlers.RegisterHandler)
	http.HandleFunc("/auth/google", oauth.GoogleLoginHandler)
	http.HandleFunc("/auth/google/callback", oauth.GoogleCallbackHandler)
	http.HandleFunc("/auth/github", oauth.GitHubLoginHandler)
	http.HandleFunc("/auth/github/callback", oauth.GitHubCallbackHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Starting server on port :8080")
	http.ListenAndServe(":8080", nil)
}
