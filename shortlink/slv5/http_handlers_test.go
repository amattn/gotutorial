package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"
)

func TestTestServer(t *testing.T) {
	res, err := http.Get(test_server.URL)
	if err != nil {
		log.Fatal(err)
	}

	assertEqual(t, 200, res.StatusCode)
}

func TestHeaderValues(t *testing.T) {
	res, err := http.Get(test_server.URL)
	if err != nil {
		log.Fatal(err)
	}

	should_be_pickles := res.Header.Get("X-MyGreatHeader")

	assertEqual(t, "Pickles", should_be_pickles)
}

func TestContextInjection(t *testing.T) {

	// users and links that exists
	test_req := httptest.NewRequest("GET", "/u/amattn/home", nil)
	_, candidate_req, _ := InjectContext(nil, test_req)
	candidate_u_value := candidate_req.Context().Value("u_value")
	assertEqual(t, "amattn", candidate_u_value)

	test_req = httptest.NewRequest("GET", "/u/frobble/ex", nil)
	_, candidate_req, _ = InjectContext(nil, test_req)
	candidate_u_value = candidate_req.Context().Value("u_value")
	assertEqual(t, "frobble", candidate_u_value)

	// users that does not exist
	test_req = httptest.NewRequest("GET", "/u/undeclared/ex", nil)
	_, candidate_req, _ = InjectContext(nil, test_req)
	candidate_u_value = candidate_req.Context().Value("u_value")
	assertEqual(t, "undeclared", candidate_u_value)

	// no user
	test_req = httptest.NewRequest("GET", "/go", nil)
	_, candidate_req, _ = InjectContext(nil, test_req)
	candidate_u_value = candidate_req.Context().Value("u_value")
	assertEqual(t, "", candidate_u_value)
}

func TestWorkingShortlinks(t *testing.T) {
	working_paths := []string{
		"go",
		"gt",
		"gh",
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

// Admin paths

func TestWorkingAdminLinks(t *testing.T) {
	not_found_paths := []string{
		"admin/new",
		"admin/list",
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
	rand.Seed(time.Now().Unix())
	random := rand.Int31()
	// so a random value will allow us to run the test multiple times, but does pollute our db.
	// normally, you'd remove the entry from the db after the test is done.
	short_code := fmt.Sprintf("test_%d", random)

	short_url := test_server.URL + "/" + short_code
	res, err := http.Get(short_url)
	if err != nil {
		log.Fatal(1819562567, short_url, err)
	}
	assertEqual(t, 404, res.StatusCode, short_url)

	// post our form
	post_url := test_server.URL + "/admin/post"
	client := http.Client{}
	form_vals := url.Values{}
	form_vals.Add("url", "http://twitter.com/GoTutorialNet")
	form_vals.Add("code", short_code)
	res, err = client.PostForm(post_url, form_vals)
	if err != nil {
		log.Fatal(1819562566, post_url, err)
	}
	assertEqual(t, 200, res.StatusCode, post_url)

	// check to make sure our short code works after
	res, err = http.Get(short_url)
	if err != nil {
		log.Fatal(1819562567, post_url, err)
	}
	assertEqual(t, 200, res.StatusCode, post_url)
}
