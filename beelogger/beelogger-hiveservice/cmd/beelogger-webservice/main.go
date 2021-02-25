package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("start beelogger-service service")
	router := NewRouter()
	log.Println("port ")
	log.Fatal(http.ListenAndServe(":8080", router))

}
