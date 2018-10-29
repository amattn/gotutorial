package main

import (
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

var show_h bool
var show_help bool
var show_version bool

func init() {
	flag.BoolVar(&show_h, "h", false, "show help message and exit(0)")
	flag.BoolVar(&show_help, "help", false, "show help message and exit(0)")
	flag.BoolVar(&show_version, "version", false, "show version info and exit(0)")
}

func main() {
	log.Println(current_function(), "entering")
	defer trace(current_function(), time.Now())

	// var err error
	log.Printf("Starting %v (v%v, build %v, build date:%v)", os.Args[0], Version(), BuildNumber(), BuildDate())
	log.Printf("os.Args: %v", os.Args)
	log.Printf("Go (runtime:%v) (GOMAXPROCS:%d) (NumCPUs:%d)\n", runtime.Version(), runtime.GOMAXPROCS(-1), runtime.NumCPU())

	// command line flags:
	flag.Parse()

	if show_version {
		os.Exit(0)
	}

	if show_h || show_help {
		flag.Usage()
		os.Exit(0)
	}

	router := NewRouter()

	host_and_port := ":8080"
	log.Println("Starting HTTP server at", "\nhttp://"+host_and_port)
	log.Fatal(http.ListenAndServe(host_and_port, router))
}

func custom404Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404 Not Found, %v", html.EscapeString(r.URL.Path))
}
