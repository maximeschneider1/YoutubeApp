package handler

import (
	"YSS/config"
	"encoding/json"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/oauth2"
	"net/http"
	"time"
)

var currentToken string

// HandleGoogleLogin builds and redirects a temporary URL to the Google consent page
// that asks for permissions for the required scopes explicitly
func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := config.GoogleOauthConfig.AuthCodeURL(config.OauthStateString)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	jsonBody, err := json.Marshal(url)
	if err != nil {
		http.Error(w, "Error converting results to json",
			http.StatusInternalServerError)
	}
	w.Write(jsonBody)

	//http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// HandleGoogleCallback is the redirect URI Oauth sends back the user to if consent is OK.
// Google POST back OAuth2 code we exchange for access token we'll use for requests to Google
func HandleGoogleCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")

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

	fmt.Println("New User has logged in")
	currentToken = token.AccessToken

	//  get user id = https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=
	// verify if user id is known (or create it) and get user data
	// put user data in jwt token
	// send back jwt as JWTcookie

	userID := "Jean Michel"

	validTokenJWT, err := generateJWT(userID, token.AccessToken)
	if err != nil {
		fmt.Println("Failed to generate token")
	}

	JWTcookie := http.Cookie{
		Name:       "jwtToken",
		Value:      validTokenJWT,
		Expires: time.Now().Add(time.Minute * 30),
	}
	http.SetCookie(w, &JWTcookie)

	loggedCookie := http.Cookie{
		Name:       "userLogged",
		Value:      "true",
		Expires: time.Now().Add(time.Minute * 30),
	}
	http.SetCookie(w, &loggedCookie)

	http.Redirect(w, r, "http://localhost:8080", http.StatusTemporaryRedirect)
}


var mySigningKey = []byte("captainjacksparrowsayshi")

func generateJWT(userID string, oauthToken string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["userID"] = userID
	claims["oauthToken"] = oauthToken
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Printf("Error signing JWT token: %v", err.Error())
		return "", err
	}

	return tokenString, nil
}

// https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=ya29.Il-9B8azP4_EYzQWXC6GPkb4jV1aGrb6k_Hsq1YfXqJB-DIcqt8gcSA8yoVucMwzmIGAqapLxsfDpC0IVsTfWV9nRN8cpwfVPsll0puO4RWBLqlN1hr6mXK0eBkxh0kPRQ


//asking  a refresh token
//POST /token HTTP/1.1
//Host: oauth2.googleapis.com
//Content-Type: application/x-www-form-urlencoded
//
//client_id=your_client_id&
//client_secret=your_client_secret&
//refresh_token=refresh_token&
//grant_type=refresh_token