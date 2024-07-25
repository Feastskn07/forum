package handlers

import (
	"forum/controller"
	"forum/render"
	"forum/session"
	"net/http"
)

var SendData = helpres.SendData{
	Data: make(map[string]interface{}), // Haritayı başlatıyoruz
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method == http.MethodPost {
		action := r.FormValue("action")
		if action == "login" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
	}

	Data := make(map[string]interface{})
	isLoggedIn, username := session.SessionCheck(r)
	Data["loggedin"] = isLoggedIn
	Data["username"] = username
	SendData.Data = Data
	render.RenderTemplate(w, "index.html", SendData.Data)
}
