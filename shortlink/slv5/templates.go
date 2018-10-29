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

type ShortlinkTemplateData struct {
	Code string
	URL  string
}

func init() {
	link_added_template = template.Must(template.New("link_added").Parse(link_added_template_source))
	list_all_template = template.Must(template.New("list_all").Parse(list_all_template_source))
}

func LinkAddedTemplateOutput(code, url string) []byte {
	buf := new(bytes.Buffer)
	data := ShortlinkTemplateData{
		Code: code,
		URL:  url,
	}
	link_added_template.Execute(buf, data)

	output, _ := ioutil.ReadAll(buf)
	return output
}

const list_all_template_source = `<html>
<body>
All shortlinks:
<br>
<table>
{{range .All}}
<tr>
<td>{{.Code}}</td>
<td>{{.URL}}</td>
</tr>
{{end}}
</table>
</body>
</html>
`

var list_all_template *template.Template

type ListAllTemplateData struct {
	All []ShortlinkTemplateData
}

func ListAllTemplateOutput(data ListAllTemplateData) []byte {
	buf := new(bytes.Buffer)
	list_all_template.Execute(buf, data)

	output, _ := ioutil.ReadAll(buf)
	return output
}
