package main

import (
	"net/http"
	"strings"
)

type LoggingRouter struct {
	adminHandler *AdminHandler
}

func NewLoggingRouter() *LoggingRouter {
	lr := new(LoggingRouter)
	lr.adminHandler = new(AdminHandler)
	return lr
}

// At a high level, a router inspects a request and routes it to an appropriate subcomponent for handling.
// Here, we just look for a simple prefix
func (router *LoggingRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	url := req.URL.Path
	if strings.HasPrefix(url, "/admin/") {
		// use the admin handler
		router.adminHandler.ServeHTTP(w, req)
	}

	// error for now
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Bad Request"))
}
