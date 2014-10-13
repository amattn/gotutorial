package main

import (
	"log"
	"net/http"
	"testing"
)

func TestTestServer(t *testing.T) {
	res, err := http.Get(test_server.URL)
	if err != nil {
		log.Fatal(err)
	}

	assertEqual(t, 200, res.StatusCode)
}

func TestWorkingShortlinks(t *testing.T) {
	working_paths := []string{
		"a",
		"b",
		"c",
	}

	for i, subpath := range working_paths {
		final_url := test_server.URL + "/" + subpath
		res, err := http.Get(final_url)
		if err != nil {
			log.Fatal(1626235976, i, subpath, final_url, err)
		}

		assertEqual(t, 200, res.StatusCode, i, final_url)
	}
}

func TestNotWorkingShortlinks(t *testing.T) {
	not_found_paths := []string{
		"bogus",
		"asdfaf",
		"1231csd",
	}

	for i, subpath := range not_found_paths {
		final_url := test_server.URL + "/" + subpath
		res, err := http.Get(final_url)
		if err != nil {
			log.Fatal(3073598627, i, subpath, err)
		}

		assertEqual(t, 404, res.StatusCode, i, final_url)
	}
}
