package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/trueaniki/admiral"
	"github.com/trueaniki/mdprinter"
)

var appName = "Markdown Printer"
var appDesc = "A simple markdown to pdf/html printer"
var version = "v0.3.0"

type Conf struct {
	Style   string `type:"flag" name:"style" alias:"s" default:"modest" description:"Style of the pdf, available styles are: modest, retro, air, splendor"`
	Align   string `type:"flag" name:"align" alias:"a" default:"left" description:"Align of the pdf content, available aligns are: left, center, right"`
	Custom  string `type:"flag" name:"custom" alias:"c" default:"none" description:"Custom css file path, if provided, it will override the style and align options"`
	Data    string `type:"flag" name:"data" alias:"d" default:"none" description:"JSON data file path, if provided, it will interpolate the data into the markdown file"`
	Output  string `type:"flag" name:"output" alias:"o" default:"none" description:"Output pdf file path"`
	Version bool   `type:"flag" name:"version" alias:"v" description:"Print version"`
	Html    bool   `type:"flag" name:"html" description:"Output html instead of pdf"`
	Input   string `type:"arg" name:"filename" required:"true" description:"Markdown file path"`
}

func main() {
	conf := &Conf{}

	a := admiral.New(appName, appDesc)
	a.Configure(conf)
	a.Flag("version").Handle(func(_ interface{}) {
		fmt.Println(appName, version, "by Aniki")
		os.Exit(0)
	})

	_, err := a.Parse(os.Args)
	if err != nil {
		printAndExit(err)
	}

	md, err := os.ReadFile(conf.Input)
	if err != nil {
		printAndExit(err)
	}
	pdfPath := getPDFPath(conf.Input)
	if conf.Output != "none" {
		pdfPath = conf.Output
	}

	p := mdprinter.New()

	if conf.Data != "none" {
		data, err := os.ReadFile(conf.Data)
		if err != nil {
			printAndExit(err)
		}
		var jsonData interface{}
		err = json.Unmarshal(data, &jsonData)
		if err != nil {
			printAndExit(err)
		}
		p.WithInterpolation(jsonData)
	}

	p.
		WithStyle(conf.Style).
		WithAlign(conf.Align).
		WithCustomCss(conf.Custom)

	if conf.Html {
		f, err := os.Create(pdfPath)
		if err != nil {
			printAndExit(err)
		}
		f.Write(p.Html(md))
		f.Close()
		return
	}

	buf, err := p.Process(md)
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
