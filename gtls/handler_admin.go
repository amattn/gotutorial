package main

import (
	"fmt"
	"html"
	"net/http"
)

type AdminHandler struct {
}

func (handler *AdminHandler) Respond(req *http.Request) (statusCode int, headers map[string]string, responseBytes []byte) {
	response := fmt.Sprintf("Hello World, you came from: %q", html.EscapeString(req.URL.Path))
	return http.StatusOK, nil, []byte(response)
}
