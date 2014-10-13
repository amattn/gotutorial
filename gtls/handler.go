package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

type LoggingHandler struct {
}

func (handler *LoggingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	output := fmt.Sprintf("Hello World, you came from: %q", html.EscapeString(r.URL.Path))
	outputBytes := []byte(output)
	w.Write(outputBytes)
	log.Printf("%s", CommonLogFormat(r, http.StatusOK, len(outputBytes)))
}
