package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Printf("start word database")
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
