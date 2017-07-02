package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func encodeResponse(w http.ResponseWriter, v interface{}) {
	err := json.NewEncoder(w).Encode(v)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		panic("Cannot encode locations")
	}
}

// Index is a root route just to Welcome! the User
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome!")
}

// LocationIndex returns all club locations
func LocationIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	locations := Locations{
		Location{Name: "North Pole"},
		Location{Name: "Equator"},
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	encodeResponse(w, locations)
}
