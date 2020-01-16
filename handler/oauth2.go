package handler

import (
	"YSS/config"
	"fmt"
	"golang.org/x/oauth2"
	"net/http"
)

var currentToken string

// HandleGoogleLogin builds and redirects a temporary URL to the Google consent page
// that asks for permissions for the required scopes explicitly
func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := config.GoogleOauthConfig.AuthCodeURL(config.OauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// HandleGoogleCallback is the redirect URI Google sends back the user to if consent is OK.
// Google POST back OAuth2 code we exchange for access token we'll use for requests to Google
func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	//w.Write([]byte(fmt.Sprintf("state is %s",state)))
	if state != config.OauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", config.OauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := config.GoogleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Println("Code exchange failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	currentToken = token.AccessToken

	fmt.Fprintf(w, htmlHome)



}