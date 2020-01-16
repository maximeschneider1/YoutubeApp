package main

import (
	"YSS/handler"
	"fmt"
	"log"
	"net/http"
)

func StartWebServer() {
	http.HandleFunc("/", handler.HandleMain)
	http.HandleFunc("/GoogleLogin", handler.HandleGoogleLogin)
	http.HandleFunc("/GoogleCallback", handler.HandleGoogleCallback)
	http.HandleFunc("/Home", handler.HandleHome)
	http.HandleFunc("/firstpage", handler.HandlePage)
	http.HandleFunc("/secondpage", handler.HandlePageTwo)
}

func main() {

	fmt.Println("Starting Web Server...")

	StartWebServer()

	log.Fatal(http.ListenAndServe(":8081", nil))

}