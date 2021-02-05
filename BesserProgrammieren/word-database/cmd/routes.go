package main

import (
	"net/http"

	"github.com/mschlech/word-database/handler"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"WordCheckResults",
		"GET",
		"wordcheck",
		handler.WordCheckResult,
	},

	Route{
		"healtCheck",
		"GET",
		"healthCheck",
		handler.HealtCheck,
	},
}
