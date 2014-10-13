package main

import (
	"log"
	"net/http"
)

const (
	LISTEN_ADDRESS = ":8080"
)

func main() {
	log.Println("gtls", Version(), "build", BuildNumber())

	logging_handler := new(LoggingHandler)
	http.Handle("/", logging_handler)

	log.Println("Listening from:", LISTEN_ADDRESS)
	log.Fatal(http.ListenAndServe(LISTEN_ADDRESS, nil))
}
