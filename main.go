package main

import (
	"log"
	"net/http"
)

type config struct {
	Port string
}

var cfg = config{
	Port: ":9000",
}

func main() {
	appHandler := WithLogging(NewRouter(routes))

	log.Printf("=> Starting server on port %s\n", cfg.Port)
	log.Fatal(http.ListenAndServe(cfg.Port, appHandler))
}
