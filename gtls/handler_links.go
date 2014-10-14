package main

import (
	"log"
	"net/http"
	"strings"
)

type LinksHandler struct {
	BaseHandler
}

func NewLinksHandler(linkstore *LinkStore) *LinksHandler {
	handler := new(LinksHandler)
	handler.BaseHandler = MakeBaseHandler(linkstore)

	handler.linkstore.AddShortlink("a", "http://golang.org")
	handler.linkstore.AddShortlink("b", "http://tour.golang.org")
	handler.linkstore.AddShortlink("c", "http://gotutorial.net")

	return handler
}

func (handler *LinksHandler) Respond(req *http.Request) (statusCode int, headers map[string]string, responseBytes []byte) {
	shortcode := req.URL.Path

	// remove leading slash if necessary
	if strings.HasPrefix(shortcode, "/") {
		shortcode = shortcode[1:]
	}

	longurl, err := handler.linkstore.GetShortlink(shortcode)
	if err != nil {
		log.Println(2097280714, err)
		return http.StatusInternalServerError, nil, []byte("Internal Server Error")
	}

	if longurl != "" {
		headers := map[string]string{"Location": longurl}
		return http.StatusMovedPermanently, headers, []byte{}
	}
	return http.StatusNotFound, nil, []byte("Not Found")
}
