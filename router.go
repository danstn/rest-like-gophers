package main

import (
	"github.com/julienschmidt/httprouter"
)

// NewRouter creates a new router and returns a pointer
func NewRouter(routes Routes) *httprouter.Router {
	router := httprouter.New()

	for _, route := range routes {
		router.Handle(route.Method, route.Path, route.Handler)
	}

	return router
}
