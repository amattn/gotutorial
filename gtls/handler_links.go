package main

import (
	"net/http"
	"strings"
)

type LinksHandler struct {
	BaseHandler
}

func NewLinksHandler(linksDB map[string]string) *LinksHandler {
	handler := new(LinksHandler)
	handler.BaseHandler = MakeBaseHandler(linksDB)

	handler.linksDB["a"] = "http://golang.org"
	handler.linksDB["b"] = "http://tour.golang.org"
	handler.linksDB["c"] = "http://gotutorial.net"

	return handler
}

func (handler *LinksHandler) Respond(req *http.Request) (statusCode int, headers map[string]string, responseBytes []byte) {
	shortcode := req.URL.Path

	// remove leading slash if necessary
	if strings.HasPrefix(shortcode, "/") {
		shortcode = shortcode[1:]
	}

	longurl := handler.linksDB[shortcode]

	if longurl != "" {
		headers := map[string]string{"Location": longurl}
		return http.StatusMovedPermanently, headers, []byte{}
	}
	return http.StatusNotFound, nil, []byte("Not Found")
}
