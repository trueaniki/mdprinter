package mdprinter_test

import (
	"strings"
	"testing"

	"github.com/trueaniki/mdprinter"
)

func TestFormupCss(t *testing.T) {
	css := mdprinter.FormupCss("air", "center", "none")
	if !strings.Contains(css, "/* air.css */") {
		t.Error("expected to form up with air.css")
	}
	if !strings.Contains(css, "text-align: center") {
		t.Error("expected to form up with align: center")
	}

	css = mdprinter.FormupCss("none", "none", "css/style/modest.css")
	if !strings.Contains(css, "/* modest.css */") {
		t.Error("expected to form up with modest.css")
	}
}
