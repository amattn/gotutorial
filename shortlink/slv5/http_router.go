package main

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"
)

type MyRouter struct {
	// do something with w and req and optionally continue
	middleware []func(http.ResponseWriter, *http.Request) (http.ResponseWriter, *http.Request, bool)
}

func NewRouter() *MyRouter {
	r := MyRouter{}

	mw := []func(http.ResponseWriter, *http.Request) (http.ResponseWriter, *http.Request, bool){
		InjectContext,
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
	current_w := w
	current_req := req
	should_continue := true

	for _, middleware := range router.middleware {
		current_w, current_req, should_continue = middleware(current_w, current_req)
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

func InjectContext(w http.ResponseWriter, req *http.Request) (http.ResponseWriter, *http.Request, bool) {
	ctx := req.Context()

	path := req.URL.Path

	u_value := ""

	if strings.HasPrefix(path, "/u/") {
		// we have a user code

		trimmed := strings.TrimPrefix(path, "/u/")
		parts := strings.SplitN(trimmed, "/", 2)
		u_value = parts[0]
	}

	updated_ctx := context.WithValue(ctx, "u_value", u_value)

	new_req := req.WithContext(updated_ctx)
	return w, new_req, true
}

func LogRequest(w http.ResponseWriter, req *http.Request) (http.ResponseWriter, *http.Request, bool) {
	log.Println(time.Now(), req)
	return w, req, true
}

// curl -I http://:8080 to see the header
func RedirectSlashes(w http.ResponseWriter, req *http.Request) (http.ResponseWriter, *http.Request, bool) {
	url := req.URL.Path

	if len(url) > 1 && strings.HasSuffix(url, "/") {
		http.Redirect(w, req, strings.TrimSuffix(url, "/"), http.StatusMovedPermanently)
		return w, req, false
	} else {
		return w, req, true
	}
}

func CustomizeHeader(w http.ResponseWriter, req *http.Request) (http.ResponseWriter, *http.Request, bool) {
	w.Header().Add("X-MyGreatHeader", "Pickles")
	return w, req, true
}
