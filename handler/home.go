package handler

import (
	"fmt"
	"net/http"
)

const htmlIndex = `<html><body><a href="/GoogleLogin">Log in with Google</a></body></html>`
const htmlHome = `<html><body><a href="/home">Go to home</a></body></html>`
const htmlPageOne = `<html><body><a href="/firstpage">Page 1</a></body></html>`
const htmlPageTwo = `<html><body><a href="/secondpage">Page 2</a></body></html>`
const htmlRandomSubReco = `<html><body><a href="/random">Random recommendation</a></body></html>`


func StartWebServer() {
	http.HandleFunc("/", HandleMain)
	http.HandleFunc("/GoogleLogin", HandleGoogleLogin)
	http.HandleFunc("/GoogleCallback", HandleGoogleCallback)
	http.HandleFunc("/home", HandleHome)
	http.HandleFunc("/firstpage", HandlePage)
	http.HandleFunc("/secondpage", HandleNextPage)
	http.HandleFunc("/random", HandleRandom)
}

//HandleMain display home html
func HandleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlIndex)
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlPageOne)
	fmt.Fprint(w, " ", htmlPageTwo)
	fmt.Fprint(w, " ", htmlRandomSubReco)
}