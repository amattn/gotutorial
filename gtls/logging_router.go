package main

import (
	"log"
	"net/http"
	"strings"
)

type LoggingRouter struct {
	adminHandler *AdminHandler
	linksHandler *LinksHandler
}

func NewLoggingRouter() *LoggingRouter {
	lr := new(LoggingRouter)
	lr.adminHandler = new(AdminHandler)
	lr.linksHandler = NewLinksHandler()
	return lr
}

type SimpleRouteHandler interface {
	Respond(req *http.Request) (statusCode int, headers map[string]string, responseBytes []byte)
}

// At a high level, a router inspects a request and routes it to an appropriate subcomponent for handling.
// Here, we just look for a simple prefix
func (router *LoggingRouter) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	url := req.URL.Path
	var code int
	responseBytes := []byte{}
	extra_headers := map[string]string{}

	if url == "/" {
		code = http.StatusOK
		responseBytes = []byte("Welcome to gtls")
	} else if strings.HasPrefix(url, "/admin/") {
		// use the admin handler
		code, extra_headers, responseBytes = router.adminHandler.Respond(req)
	} else {
		// use the shortlink handler
		code, extra_headers, responseBytes = router.linksHandler.Respond(req)
	}

	for k, v := range extra_headers {
		w.Header().Add(k, v)
	}

	w.WriteHeader(code)
	writtenCount, err := w.Write(responseBytes)
	if err != nil {
		log.Println("error writing response", req, err)
	}
	log.Printf("%s", CommonLogFormat(req, code, writtenCount))
}
