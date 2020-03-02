package handler

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	//"github.com/julienschmidt/httprouter"
)

const htmlIndex = `<html><body><a href="/GoogleLogin">Log in with Google</a></body></html>`
const htmlHome = `<html><body><a href="/home">Go to home</a></body></html>`
const htmlPageOne = `<html><body><a href="/firstpage">Page 1</a></body></html>`
const htmlPageTwo = `<html><body><a href="/secondpage">Page 2</a></body></html>`
const htmlRandomSubReco = `<html><body><a href="/random">Random recommendation</a></body></html>`

//func handleNext(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
//	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
//}
func handleSomething() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// use thing
	}
}
func StartWebServer() {
	//
	router := httprouter.New()
	//router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	if r.Header.Get("Access-Control-Request-Method") != "" {
	//		// Set CORS headers
	//		header := w.Header()
	//		header.Set("Access-Control-Allow-Methods", r.Header.Get("Allow"))
	//		header.Set("Access-Control-Allow-Origin", "*")
	//	}
	//})
	//
	//router.POST("/nextsub", handleNext)
	router.HandlerFunc("GET", "/api", handleSomething())

	http.HandleFunc("/", HandleMain)
	http.HandleFunc("/GoogleLogin", HandleGoogleLogin)
	http.HandleFunc("/GoogleCallback", HandleGoogleCallback)
	http.HandleFunc("/home", HandleHome)
	http.HandleFunc("/subscriptions", HandleGetSubs)
	http.HandleFunc("/nextsubs", HandleNextSubs)
	http.HandleFunc("/random", random)
	http.HandleFunc("/get", HandleGet)

	log.Fatal(http.ListenAndServe(":8081", nil))
}
func HandleGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	results := "Hello Maxime"
	jsonBody, err := json.Marshal(results)
	if err != nil {
		http.Error(w, "Error converting results to json",
			http.StatusInternalServerError)
	}
	w.Write(jsonBody)
}

//HandleMain display home html
func HandleMain(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)

	fmt.Fprintf(w, htmlIndex)
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, htmlPageOne)
	fmt.Fprint(w, " ", htmlPageTwo)
	fmt.Fprint(w, " ", htmlRandomSubReco)
}