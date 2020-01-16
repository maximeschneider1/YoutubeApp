package main

import (
	"YSS/handler"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", handler.HandleMain)
	http.HandleFunc("/GoogleLogin", handler.HandleGoogleLogin)
	http.HandleFunc("/GoogleCallback", handler.HandleGoogleCallback)
	http.HandleFunc("/Home", handler.HandleHome)
	http.HandleFunc("/firstpage", handler.HandlePageOne)
	http.HandleFunc("/secondpage", handler.HandlePageTwo)

	log.Fatal(http.ListenAndServe(":8081", nil))











}




