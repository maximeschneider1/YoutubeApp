package handler

import (
	"fmt"
	"net/http"
)
const htmlIndex = `<html><body><a href="/GoogleLogin">Log in with Google</a></body></html>`
const htmlHome = `<html><body><a href="/Home">Go to home</a></body></html>`
const htmlPageOne = `<html><body><a href="/firstpage">Page 1</a></body></html>`
const htmlPageTwo = `<html><body><a href="/secondpage">Page 2</a></body></html>`


//HandleMain display home html
func HandleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlIndex)
}

func HandleHome(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, htmlPageOne)

	fmt.Fprint(w, " ", htmlPageTwo)

}