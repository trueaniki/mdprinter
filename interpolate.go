package mdprinter

import (
	"bytes"
	"text/template"
)

func Interpolate(html []byte, data interface{}) []byte {
	template, err := template.New("template").Parse(string(html))
	if err != nil {
		panic(err)
	}
	var buf bytes.Buffer
	template.Execute(&buf, data)
	return buf.Bytes()
}
