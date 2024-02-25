package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/jessevdk/go-flags"
	"github.com/trueaniki/mdprinter"
)

var version = "0.0.0"

var opts struct {
	Style       string      `short:"s" long:"style" description:"style of the pdf, available styles are: modest, retro, air, splendor" default:"modest"`
	Align       string      `short:"a" long:"align" description:"align of the pdf content, available aligns are: left, center, right" default:"left"`
	Custom      string      `short:"c" long:"custom" description:"custom css file path, if provided, it will override the style and align options" default:"none"`
	Data        string      `short:"d" long:"data" description:"JSON data file path, if provided, it will interpolate the data into the markdown file" default:"none"`
	Output      string      `short:"o" long:"output" description:"output pdf file path" default:"none"`
	Version     bool        `short:"v" long:"version" description:"print version"`
	Positionals positionals `positional-args:"filename" required:"true"`
}

type positionals struct {
	Filename string `positional-arg-name:"filename" required:"true"`
}

func main() {
	args, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		printAndExit(err)
	}
	if opts.Version {
		fmt.Println(version)
		os.Exit(0)
	}

	mdPath := args[1]

	md, err := os.ReadFile(mdPath)
	if err != nil {
		printAndExit(err)
	}
	pdfPath := getPDFPath(mdPath)
	if opts.Output != "none" {
		pdfPath = opts.Output
	}

	html := mdprinter.Parse(md)

	if opts.Data != "none" {
		data, err := os.ReadFile(opts.Data)
		if err != nil {
			printAndExit(err)
		}
		var jsonData interface{}
		err = json.Unmarshal(data, &jsonData)
		if err != nil {
			printAndExit(err)
		}
		html = mdprinter.Interpolate(html, jsonData)
	}

	buf, err := mdprinter.Print(html, mdprinter.FormupCss(opts.Style, opts.Align, opts.Custom))
	if err != nil {
		printAndExit(err)
	}

	f, err := os.Create(pdfPath)
	if err != nil {
		printAndExit(err)
	}

	f.Write(buf)
	f.Close()
}

func getPDFPath(p string) string {
	withoutExt := path.Base(p)[:len(path.Base(p))-len(path.Ext(p))]
	return path.Join(path.Dir(p), withoutExt+".pdf")
}

func printAndExit(err error) {
	fmt.Println(err)
	os.Exit(1)
}
