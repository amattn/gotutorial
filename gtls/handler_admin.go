package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

type AdminHandler struct {
}

func (handler *AdminHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	output := fmt.Sprintf("Hello World, you came from: %q", html.EscapeString(req.URL.Path))
	outputBytes := []byte(output)
	w.Write(outputBytes)
	log.Printf("%s", CommonLogFormat(req, http.StatusOK, len(outputBytes)))
}
