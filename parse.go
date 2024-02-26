package mdprinter

import (
	"bytes"

	"github.com/yuin/goldmark"
)

func Parse(md []byte) []byte {
	p := goldmark.New()
	var buf bytes.Buffer
	if err := p.Convert(md, &buf); err != nil {
		panic(err)
	}
	return buf.Bytes()
}
