# Markdown Printer
Prints markdown files to PDF.
## CLI Usage
`mdprinter test.md -s air` outputs PDF file `test.pdf`.
## API Usage
```go
p := mdprinter.New()
pdfBuf, err := p.
  WithInterpolation(struct{
    Name string
	  Age string
  }{
	  Name: "John",
	  Age: "25",
  }).
  WithStyle("air").
  WithAlign("left").
  WithCustomCss("./mycustom.css").
  Process(markdownData)
```
## Styles
### Built-in styles
4 styles from [https://markdowncss.github.io/](https://markdowncss.github.io/) available:
- *modest*
- *retro*
- *air*
- *splendor*
Use them with `-s` or `--style` flag.
### Alignment
You can set contents alignment using `-a` or `--align` flag. Available options are:
- *left*
- *center*
- *right*
### Custom styles
You can provide your own custom css using `-c` or `--custom` flag. Set it to the path of your CSS file.
## Interpolation
You can interpolate data into markdown file. To do this, you need to have your data in JSON file. In the markdown file follow the rules of templating from [https://pkg.go.dev/text/template](pkg.go.dev/text/template). Then run programm with `-d` or `--data` flag set to the path of your file.

## Examples
Play with files in `examples/` folder. Run 

```sh
go run cmd/mdprinter/main.go examples/example.md -s retro -a left -d examples/exampleData.json
``` 

to get the PDF.