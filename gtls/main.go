package main

import "log"

func main() {
	log.Println("gtls", Version(), "build", BuildNumber())

	log.Println("Hello World")
}
