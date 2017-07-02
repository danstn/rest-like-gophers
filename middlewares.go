package main

import (
	"log"
	"net/http"
)

type loggingHandler struct {
	handler http.Handler
}

func (h loggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
	log.Printf("%s %s %s", r.Method, r.RemoteAddr, r.URL)
}

// WithLogging decorates the given handler with basic logging
func WithLogging(handler http.Handler) http.Handler {
	return loggingHandler{handler: handler}
}
