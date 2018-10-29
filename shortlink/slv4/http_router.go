package main

import (
	"log"
	"net/http"
	"strings"
	"time"
)

type MyRouter struct {
	// do something with w and req and optionally continue
	middleware []func(http.ResponseWriter, *http.Request) bool
}

func NewRouter() *MyRouter {
	r := MyRouter{}

	mw := []func(http.ResponseWriter, *http.Request) bool{
		LogRequest,
		RedirectSlashes,
		CustomizeHeader,
		//Throttle
		//Compress
		//etc.
	}

	r.middleware = mw
	return &r
}

// At a high level, a router inspects a request and routes it to an appropriate subcomponent for handling.
// Here, we just look for a simple prefix
func (router *MyRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	url := req.URL.Path

	// middleware is essentially just a function we call on every incoming request
	for _, middleware := range router.middleware {
		should_continue := middleware(w, req)
		if should_continue == false {
			return
		}
	}

	switch {
	case url == "/":
		rootHandler(w, req)
	case url == "/admin/new":
		newShortlinkHandler(w, req)
	case url == "/admin/post":
		postHandler(w, req)
	case url == "/admin/success":
		successHandler(w, req)
	case url == "/admin/list":
		listAllHandler(w, req)
	default:
		shortlinkHandler(w, req)
	}
}

func LogRequest(w http.ResponseWriter, req *http.Request) bool {
	log.Println(time.Now(), req)
	return true
}

// curl -I http://:8080 to see the header
func RedirectSlashes(w http.ResponseWriter, req *http.Request) bool {
	url := req.URL.Path

	if len(url) > 1 && strings.HasSuffix(url, "/") {
		http.Redirect(w, req, strings.TrimSuffix(url, "/"), http.StatusMovedPermanently)
		return false
	} else {
		return true
	}
}

func CustomizeHeader(w http.ResponseWriter, req *http.Request) bool {
	w.Header().Add("X-MyGreatHeader", "Pickles")
	return true
}
