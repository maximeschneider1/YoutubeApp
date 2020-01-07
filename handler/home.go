package handler

import (
	"fmt"
	"net/http"
)
const htmlIndex = `<html><body><a href="/GoogleLogin">Log in with Google</a></body></html>`

//HandleMain display home html
func HandleMain(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlIndex)
}