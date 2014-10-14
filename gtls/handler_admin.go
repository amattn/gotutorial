package main

import (
	"fmt"
	"html"
	"net/http"
)

type AdminHandler struct {
	BaseHandler
}

func NewAdminHandler(linksDB map[string]string) *AdminHandler {
	handler := new(AdminHandler)
	handler.BaseHandler = MakeBaseHandler(linksDB)
	return handler
}

func (handler *AdminHandler) Respond(req *http.Request) (statusCode int, headers map[string]string, responseBytes []byte) {
	url_path := req.URL.Path
	switch url_path {
	case "/admin/post":
		if req.Method == "POST" {
			// temporary just respond okay
			return http.StatusOK, nil, []byte("ok")
		} else {
			return http.StatusMethodNotAllowed, nil, []byte("Method not allowed")
		}
	}

	response := fmt.Sprintf("Hello World, you came from: %q", html.EscapeString(req.URL.Path))
	return http.StatusOK, nil, []byte(response)
}
