package main

import (
	"log"
	"net/http"
	"testing"
)

func TestWorkingAdminLinks(t *testing.T) {
	not_found_paths := []string{
		"admin/",
		"admin/bogus",
		"admin/asdfaf",
	}

	for i, subpath := range not_found_paths {
		final_url := test_server.URL + "/" + subpath
		res, err := http.Get(final_url)
		if err != nil {
			log.Fatal(3073598627, i, subpath, err)
		}

		assertEqual(t, 200, res.StatusCode, i, final_url)
	}
}
