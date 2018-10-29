package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
	"log"
)

// best practice:
// use a package like packr (https://github.com/gobuffalo/packr)
// to embed .html from your source tree rather than use const in source code.

const link_added_template_source = `<html>
<body>
You can find the url:<a href="{{.Data.URL}}">{{.Data.URL}}</a>
at with the <a href="/{{.Data.Code}}">shortcode {{.Data.Code}}</a>
</body>
</html>
`
const list_all_template_source = `<html>
<head>
	<title>{{.Title}}</title>
</head>
<body>
<h1>{{.Title}}</h1>
<br>
<table>
{{ range $key, $value := .Data.All }}
<tr>
<td>{{$key}}</td>
<td>{{$value}}</td>
</tr>
{{end}}
</table>
</body>
</html>
`

var link_added_template *template.Template
var list_all_template *template.Template

// best practice (for dev efficiency):
// utilize a flexible, reusable template data structure rather
// than a custom data structure for each template.
type ContentData struct {
	Title  string
	Data   map[string]interface{} // any data the page may need
	Errors map[string]interface{} // misc form validation errors, etc.
}

func init() {

	// compile our templates

	link_added_template = template.Must(template.New("link_added").Parse(link_added_template_source))
	list_all_template = template.Must(template.New("list_all").Parse(list_all_template_source))
}

func LinkAddedTemplateOutput(code, url string) []byte {

	data := ContentData{
		Title: "Link Added",
		Data: map[string]interface{}{
			"Code": code,
			"URL":  url,
		},
		// no errors
	}

	buf := new(bytes.Buffer)
	link_added_template.Execute(buf, data)

	output, _ := ioutil.ReadAll(buf)
	return output
}

func ListAllTemplateOutput(shortlinks map[string]string) []byte {

	data := ContentData{
		Title: "All Shortlinks",
		Data: map[string]interface{}{
			"All": shortlinks,
		},
		// no errors
	}

	log.Println("3300263639, ", data)

	buf := new(bytes.Buffer)
	list_all_template.Execute(buf, data)

	output, _ := ioutil.ReadAll(buf)
	return output
}
