package mdprinter_test

import (
	"encoding/json"
	"testing"

	"github.com/trueaniki/mdprinter"
)

func TestInterpolate(t *testing.T) {
	template := []byte("{{.Name}} is {{.Age}} years old")
	jsonData := "{\"Name\": \"John\", \"Age\": 30}"
	var data interface{}
	json.Unmarshal([]byte(jsonData), &data)
	got := mdprinter.Interpolate(template, data)
	want := []byte("John is 30 years old")
	if string(got) != string(want) {
		t.Errorf("got \"%s\", want \"%s\"", got, want)
	}
}
