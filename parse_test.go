package mdprinter_test

import (
	"strings"
	"testing"

	"github.com/trueaniki/mdprinter"
)

func TestParse(t *testing.T) {
	for _, test := range tests {
		got := string(mdprinter.Parse([]byte(test.input)))
		got = strings.ReplaceAll(got, "\n", "")
		if got != test.output {
			t.Errorf("got %s, want %s", got, test.output)
		}
	}
}

var tests = []struct {
	input  string
	output string
}{
	{"", ""},
	{"# Hello", "<h1>Hello</h1>"},
	{"## Hello", "<h2>Hello</h2>"},
	{"### Hello", "<h3>Hello</h3>"},
	{"#### Hello", "<h4>Hello</h4>"},
	{"##### Hello", "<h5>Hello</h5>"},
	{"###### Hello", "<h6>Hello</h6>"},
	{"Hello", "<p>Hello</p>"},
	{"Hello\nWorld", "<p>Hello<br>World</p>"},
	{"Hello\n\nWorld", "<p>Hello</p><p>World</p>"},
	{"Hello\n\nWorld\n\n", "<p>Hello</p><p>World</p>"},
	{"Hello\n\nWorld\n\n\n", "<p>Hello</p><p>World</p>"},
	{"Hello\n\nWorld\n\n\n\n", "<p>Hello</p><p>World</p>"},
	{"Hello\n\nWorld\n\n\n\n\n", "<p>Hello</p><p>World</p>"},
	{"Hello\n\nWorld\n\n\n\n\n\n", "<p>Hello</p><p>World</p>"},
	{"Hello\n\nWorld\n\n\n\n\n\n\n", "<p>Hello</p><p>World</p>"},
	{"Hello\n\nWorld\n\n\n\n\n\n\n\n", "<p>Hello</p><p>World</p>"},
	{"Hello\n\nWorld\n\n\n\n\n\n\n\n", "<p>Hello</p><p>World</p>"},
	{"*Hello*", "<em>Hello</em>"},
	{"_Hello_", "<em>Hello</em>"},
	{"**Hello**", "<strong>Hello</strong>"},
	{"__Hello__", "<strong>Hello</strong>"},
	{"***Hello***", "<strong><em>Hello</em></strong>"},
	{"___Hello___", "<strong><em>Hello</em></strong>"},
	{"~~Hello~~", "<del>Hello</del>"},
	{"`Hello`", "<code>Hello</code>"},
	{"```Hello```", "<pre><code>Hello</code></pre>"},
	{"[Hello](https://example.com)", "<a href=\"https://example.com\">Hello</a>"},
	{"![Hello](https://example.com)", "<img src=\"https://example.com\" alt=\"Hello\">"},
	{"- Hello", "<ul><li>Hello</li></ul>"},
	{"- Hello\n- World", "<ul><li>Hello</li><li>World</li></ul>"},
	{"1. Hello", "<ol><li>Hello</li></ol>"},
	{"1. Hello\n2. World", "<ol><li>Hello</li><li>World</li></ol>"},
	{"- [ ] Hello", "<ul><li><input type=\"checkbox\">Hello</li></ul>"},
	{"- [x] Hello", "<ul><li><input type=\"checkbox\" checked>Hello</li></ul>"},
	{"- [X] Hello", "<ul><li><input type=\"checkbox\" checked>Hello</li></ul>"},
	{"- [x] Hello\n- [ ] World", "<ul><li><input type=\"checkbox\" checked>Hello</li><li><input type=\"checkbox\">World</li></ul>"},
	{"- [x] Hello\n- [x] World", "<ul><li><input type=\"checkbox\" checked>Hello</li><li><input type=\"checkbox\" checked>World</li></ul>"},
	{"- [x] Hello\n- [x] World\n- [ ] Universe", "<ul><li><input type=\"checkbox\" checked>Hello</li><li><input type=\"checkbox\" checked>World</li><li><input type=\"checkbox\">Universe</li></ul>"},
	{"- [x] Hello\n- [x] World\n- [x] Universe", "<ul><li><input type=\"checkbox\" checked>Hello</li><li><input type=\"checkbox\" checked>World</li><li><input type=\"checkbox\" checked>Universe</li></ul>"},
	{"- [x] Hello\n- [x] World\n- [x] Universe\n- [ ] Galaxy", "<ul><li><input type=\"checkbox\" checked>Hello</li><li><input type=\"checkbox\" checked>World</li><li><input type=\"checkbox\" checked>Universe</li><li><input type=\"checkbox\">Galaxy</li></ul>"},
	{"- [x] Hello\n- [x] World\n- [x] Universe\n- [x] Galaxy", "<ul><li><input type=\"checkbox\" checked>Hello</li><li><input type=\"checkbox\" checked>World</li><li><input type=\"checkbox\" checked>Universe</li><li><input type=\"checkbox\" checked>Galaxy</li></ul>"},
}
