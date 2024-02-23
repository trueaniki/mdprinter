package mdprinter

import (
	"embed"
	"os"
)

//go:embed css/*
var css embed.FS

func FormupCss(style string, align string, custom string) string {
	var css string
	css += getCss("style", style)
	css += getCss("align", align)
	if custom != "none" {
		css += getCss("custom", custom)
	}
	return css
}

func getCss(folder string, filename string) string {
	if filename == "none" || filename == "" {
		return ""
	}
	if folder == "custom" {
		data, err := os.ReadFile(filename)
		if err != nil {
			panic(err)
		}
		return string(data)
	}
	data, err := css.ReadFile("css/" + folder + "/" + filename + ".css")
	if err != nil {
		panic(err)
	}
	return string(data)
}
