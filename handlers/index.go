package handlers

import (
	"fmt"
	"forum/connection"
	"forum/helpers"
	"forum/render"
	auth "forum/session"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	Data := make(map[string]interface{})

	Data["loggedin"] = auth.SessionCheck(r)

	posts, err := connection.GetThreePosts(connection.DB)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cookie, err := r.Cookie("session_token")
	if cookie != nil && err == nil {
		sessionToken := cookie.Value
		username := auth.Sessions[sessionToken].Username

		Data["user"], err = connection.GetUser(connection.DB, username)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		Data["userinfo"], _ = connection.GetProfile(connection.DB, helpers.CreateUrl(username))
	}

	Data["posts"] = posts
	helpers.SendData.Data = Data
	render.RenderTemplate(w, "index-page.html", helpers.SendData.Data)
}
