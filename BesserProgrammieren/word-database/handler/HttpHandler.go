package handler

import (
	"log"
	"net/http"
)

func WordCheckResult(w http.ResponseWriter, r *http.Request) {

	jsonResponse(w, http.StatusOK, "test")
	return
}

func HealtCheck(w http.ResponseWriter , r *http.Request) {

	log.Printf("healtCheck" , r)
}
