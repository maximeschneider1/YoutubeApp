package config

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"math/rand"
	"os"
)

type (
	JsonFile struct {
		Web myGoogleConfig `json:"web"`
	}
	myGoogleConfig struct {
		ClientID   string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	}
)

var (
	GoogleOauthConfig = &oauth2.Config{
		RedirectURL:    "http://localhost:8081/GoogleCallback",
		ClientID:     getConfig().Web.ClientID,
		ClientSecret: getConfig().Web.ClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/youtube.readonly"},
		Endpoint:     google.Endpoint,
	}
	googleSecretPath = "config/code_secret_client_139508597776-6903ekrtdtc041j9eebtsumcl2jglmf6.apps.googleusercontent.com.json"
	// Some random string, random for each request
	OauthStateString = randState(7)
)


func randState(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}



func getConfig() JsonFile {

	jsonFile, err := os.Open(googleSecretPath); if err != nil {
		fmt.Println(err.Error())
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var GoogleConfig = JsonFile{}
	json.Unmarshal(byteValue, &GoogleConfig)

	return GoogleConfig
}
