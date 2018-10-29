package main

import (
	"net/http"
)

type MyRouter struct {
}

func NewRouter() *MyRouter {
	return new(MyRouter)
}

// At a high level, a router inspects a request and routes it to an appropriate subcomponent for handling.
// Here, we just look for a simple prefix
func (router *MyRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	url := req.URL.Path

	switch {
	case url == "/":
		rootHandler(w, req)
	default:
		shortlinkHandler(w, req)
	}
}
