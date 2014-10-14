package main

import "net/http"

type AdminHandler struct {
	BaseHandler
}

func NewAdminHandler(linkstore *LinkStore) *AdminHandler {
	handler := new(AdminHandler)
	handler.BaseHandler = MakeBaseHandler(linkstore)
	return handler
}

func (handler *AdminHandler) ListAllShortlinks(req *http.Request) (statusCode int, headers map[string]string, responseBytes []byte) {
	if req.Method != "GET" {
		return http.StatusMethodNotAllowed, nil, []byte("Method not allowed")
	}

	all_shortlinks, err := handler.linkstore.GetAllShortlinks()
	if err != nil {
		return http.StatusInternalServerError, nil, []byte("Internal Server Error")
	}

	output := ListAllTemplateOutput(ListAllTemplateData{all_shortlinks})
	return http.StatusOK, nil, output
}

func (handler *AdminHandler) AddShortlinkFormResponse(req *http.Request) (statusCode int, headers map[string]string, responseBytes []byte) {
	if req.Method != "GET" {
		return http.StatusMethodNotAllowed, nil, []byte("Method not allowed")
	}
	form_html := `<html>
<form action="/admin/post" method="POST">
URL: <input type="text" name="url"><br>
Short Code: <input type="text" name="code"><br>
<input type="submit" value="Submit">
</form>
</html>
`
	return http.StatusOK, nil, []byte(form_html)
}

func (handler *AdminHandler) PostResponse(req *http.Request) (statusCode int, headers map[string]string, responseBytes []byte) {
	if req.Method != "POST" {
		return http.StatusMethodNotAllowed, nil, []byte("Method not allowed")
	}

	url_path := req.URL.Path
	switch url_path {
	case "/admin/post":
		code := req.FormValue("code")
		url := req.FormValue("url")
		handler.linkstore.AddShortlink(code, url)
		output := LinkAddedTemplateOutput(code, url)
		// normally you want to redirect instead of returning direction...
		// otherwise the user can reload and unintentionally post the same data multiple times.
		return http.StatusOK, nil, output
	default:
		return http.StatusMethodNotAllowed, nil, []byte("Method not allowed")
	}
}
