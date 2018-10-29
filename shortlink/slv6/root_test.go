package main

import (
	"log"
	"net/http"
	"testing"
)

func TestRoot(t *testing.T) {
	url := test_server.URL
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	assertEqual(t, 200, res.StatusCode, url)
}

func Test404(t *testing.T) {
	url := test_server.URL + "/NOT_A_REAL_URL/xlkjchfa9"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	assertEqual(t, 404, res.StatusCode, url)
}
