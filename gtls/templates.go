package main

import (
	"bytes"
	"html/template"
	"io/ioutil"
)

const link_added_template_source = `<html>
<body>
You can find the url:<a href="{{.URL}}">{{.URL}}</a>
at with the <a href="/{{.Code}}">shortcode {{.Code}}</a>
</body>
</html>
`

var link_added_template *template.Template

type TemplateData struct {
	URL  string
	Code string
}

func init() {
	link_added_template = template.Must(template.New("link_added").Parse(link_added_template_source))
}

func LinkAddedTemplateOutput(code, url string) []byte {
	buf := new(bytes.Buffer)
	data := TemplateData{
		URL:  url,
		Code: code,
	}
	link_added_template.Execute(buf, data)

	output, _ := ioutil.ReadAll(buf)
	return output
}
