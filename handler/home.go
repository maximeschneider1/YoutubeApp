package handler

import (
	"YoutubeApp/config"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
)


type server struct {
	db *sql.DB
	router *httprouter.Router
}

const htmlIndex = `<html><body><a href="/GoogleLogin">Log in with Google</a></body></html>`
const htmlHome = `<html><body><a href="/home">Go to home</a></body></html>`
const htmlPageOne = `<html><body><a href="/firstpage">Page 1</a></body></html>`
const htmlPageTwo = `<html><body><a href="/secondpage">Page 2</a></body></html>`
const htmlRandomSubReco = `<html><body><a href="/random">Random recommendation</a></body></html>`

var configPath = "/Users/max/go/src/YoutubeApp/config/DBconfig.json"

func StartWebServer() {

	db, err := config.ReturnDB(configPath); if err != nil {
		fmt.Println(err.Error())
	}

	s := server{
		db: db,
		router: httprouter.New(),
	}
	s.router.GlobalOPTIONS = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			header := w.Header()
			header.Set("Access-Control-Allow-Methods", r.Header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
			header.Set("Access-Control-Allow-Headers", "*")
		}
		w.WriteHeader(http.StatusNoContent)
	})

	s.routes()

	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), s.router))

	//http.HandleFunc("/GoogleLogin", HandleGoogleLogin)
	//http.HandleFunc("/GoogleCallback", HandleGoogleCallback)
	//http.HandleFunc("/home", HandleHome)
	//http.HandleFunc("/subscriptions", HandleGetSubs)
	//http.HandleFunc("/nextsubs", HandleNextSubs)
	//http.HandleFunc("/random", random)
}

func (s *server) routes() {
	s.router.HandlerFunc("GET", "/api", s.handleSomething())
	s.router.HandlerFunc("GET", "/get", s.HandleGet())
	s.router.HandlerFunc("GET", "/", s.HandleMain())

	s.router.HandlerFunc("GET", "/GoogleLogin", s.HandleGoogleLogin())
	s.router.HandlerFunc("GET", "/GoogleCallback", s.HandleGoogleCallback())

	s.router.HandlerFunc("GET", "/home", s.HandleHome())
	s.router.HandlerFunc("GET", "/subscriptions", s.HandleGetSubs())
	s.router.HandlerFunc("GET", "/nextsubs", s.HandleNextSubs())
	s.router.HandlerFunc("GET", "/random", s.random())
}

func (s *server)  handleSomething() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// use thing
	}
}

func (s *server)  HandleGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
}

//HandleMain display home html
func (s *server)  HandleMain() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprintf(w, htmlIndex)
	}
}

func (s *server)HandleHome() http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, htmlPageOne)
		fmt.Fprint(w, " ", htmlPageTwo)
		fmt.Fprint(w, " ", htmlRandomSubReco)
	}
}