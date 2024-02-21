package main

import (
	"fmt"
	"os"
	"path"

	"github.com/jessevdk/go-flags"
	"github.com/trueaniki/mdprinter"
)

var opts struct {
	Style string `short:"s" long:"style" description:"style of the pdf, available styles are: modest, retro, air, splendor" default:"modest"`
}

func main() {
	args, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	mdPath := args[1]

	md, err := os.ReadFile(mdPath)
	if err != nil {
		panic(err)
	}
	pdfPath := getPDFPath(mdPath)

	html := mdprinter.Parse(md)

	buf, err := mdprinter.Print(html, opts.Style)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(pdfPath)
	if err != nil {
		panic(err)
	}

	f.Write(buf)
	f.Close()
}

func getPDFPath(p string) string {
	withoutExt := path.Base(p)[:len(path.Base(p))-len(path.Ext(p))]
	return path.Join(path.Dir(p), withoutExt+".pdf")
}
