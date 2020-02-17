package main

import (
	"YSS/handler"
	"fmt"
	"log"
	"net/http"
)



func main() {


	fmt.Println("Starting Web Server...")

	handler.StartWebServer()

	log.Fatal(http.ListenAndServe(":8081", nil))

}