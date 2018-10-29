package main

import (
	"flag"
	"log"
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

	log.Println("Hello Gophers!")
}
