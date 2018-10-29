package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"
)

var shortlinks map[string]string

func init() {
	shortlinks = make(map[string]string)

	shortlinks["go"] = "http://golang.org"
	shortlinks["gt"] = "http://gotutorial.net"
	shortlinks["gh"] = "http://github.com"
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "root!")
}

func shortlinkHandler(w http.ResponseWriter, req *http.Request) {
	url := req.URL.Path

	shortcode := strings.TrimPrefix(url, "/")

	dest_url, shortcode_exists := shortlinks[shortcode]

	log.Println(shortcode, shortcode_exists, dest_url)

	if shortcode_exists {
		http.Redirect(w, req, dest_url, http.StatusMovedPermanently)
		return
	} else {
		my404Handler(w, req)
		return
	}
}

func my404Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404 Not Found, %v", html.EscapeString(r.URL.Path))
}
