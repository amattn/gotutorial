package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"
)

var shortlinks map[string]string

func init() {
	shortlinks = make(map[string]string)

	shortlinks["go"] = "http://golang.org"
	shortlinks["gt"] = "http://gotutorial.net"
	shortlinks["gh"] = "http://github.com"

	shortlinks["u/amattn/g"] = "http://golang.org"
	shortlinks["u/amattn/home"] = "http://amattn.com"

	shortlinks["u/frobble/ex"] = "http://example.com"
}

func addShortlink(shortcode, url string) {
	shortlinks[shortcode] = url
}

func rootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "root!")
}

func newShortlinkHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
	form_html := `<html>
<form action="/admin/post" method="POST">
URL: <input type="text" name="url"><br>
Short Code: <input type="text" name="code"><br>
<input type="submit" value="Submit">
</form>
</html>
`
	fmt.Fprintf(w, form_html)

	return
}

func postHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	url_path := req.URL.Path
	switch url_path {
	case "/admin/post":
		shortcode := req.FormValue("code")
		url := req.FormValue("url")
		addShortlink(shortcode, url)

		// normally you want to redirect instead of returning direction...
		// otherwise the user can reload and unintentionally post the same data multiple times.
		http.Redirect(w, req, "/admin/success", http.StatusSeeOther)
		return
	default:
		my404Handler(w, req)
		return
	}
}

func successHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "success!")
}

func listAllHandler(w http.ResponseWriter, req *http.Request) {

	output := ListAllTemplateOutput(shortlinks)

	// old
	// fmt.Fprintf(w, "<html>")
	// fmt.Fprintf(w, "</html>")

	// new
	_, _ = w.Write(output)
	// please don't ignore errors.  prob should log something here if w.Write errors for some reason.
}

func shortlinkHandler(w http.ResponseWriter, req *http.Request) {
	url := req.URL.Path

	shortcode := strings.TrimPrefix(url, "/")

	dest_url, shortcode_exists := shortlinks[shortcode]

	log.Println(shortcode, shortcode_exists, dest_url)

	if shortcode_exists {
		http.Redirect(w, req, dest_url, http.StatusMovedPermanently)
		return
	} else {
		my404Handler(w, req)
		return
	}
}

func my404Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "404 Not Found, %v", html.EscapeString(r.URL.Path))
}
