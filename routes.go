package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Route represents a net/http compatible route
type Route struct {
	Name    string
	Path    string
	Method  string
	Handler httprouter.Handle
}

// Routes a collection of routes
type Routes []Route

var routes = Routes{
	Route{
		Name:    "Index",
		Path:    "/",
		Method:  http.MethodGet,
		Handler: Index,
	},
	Route{
		Name:    "Locations",
		Path:    "/locations",
		Method:  http.MethodGet,
		Handler: LocationIndex,
	},
}
