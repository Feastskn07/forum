package oauth

import (
	"context"
	"forum/handlers"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	goauth2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

var googleOAuthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8080/auth/google/callback",
	ClientID:     "659027315220-q98ejk1g2sk4k2659v1g9t1h81fcufp0.apps.googleusercontent.com",
	ClientSecret: "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",
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

func GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := googleOAuthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	handlers.HandleOAuthCallback(w, r, googleOAuthConfig, getGoogleUserInfo)
}

func getGoogleUserInfo(token *oauth2.Token) (string, string, error) {
	client := googleOAuthConfig.Client(context.Background(), token)
	oauth2Service, err := goauth2.NewService(context.Background(), option.WithHTTPClient(client))
	if err != nil {
		return "", "", err
	}
	userInfoService := goauth2.NewUserinfoService(oauth2Service)
	userInfo, err := userInfoService.Get().Do()
	if err != nil {
		return "", "", err
	}

	return userInfo.Name, userInfo.Email, nil
}
