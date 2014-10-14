package main

import (
	"log"
	"net/http"
	"net/url"
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
			log.Fatal(3377573872, i, subpath, err)
		}

		assertEqual(t, 200, res.StatusCode, i, final_url)
	}
}

func TestGetAPostRoute(t *testing.T) {
	post_url := test_server.URL + "/admin/post"
	res, err := http.Get(post_url)
	if err != nil {
		log.Fatal(1207235306, post_url, err)
	}

	assertEqual(t, 405, res.StatusCode, post_url)
}

func TestCreateShortlink(t *testing.T) {
	post_url := test_server.URL + "/admin/post"

	client := http.Client{}
	form_vals := url.Values{}
	form_vals.Add("url", "http://twitter.com/GoTutorialNet")
	res, err := client.PostForm(post_url, form_vals)
	if err != nil {
		log.Fatal(1819562566, post_url, err)
	}

	assertEqual(t, 200, res.StatusCode, post_url)

}