# Markdown Printer
Prints markdown files to PDF
## Usage
`mdprinter test.md -s air` outputs PDF file `test.pdf`
## Styles
4 styles from [https://markdowncss.github.io/](https://markdowncss.github.io/) available:
- *modest*
- *retro*
- *air*
- *splendor*


Use them with `-s` or `--style` flag

## Interpolation
You can interpolate data into markdown file. To do this, you need to have your data in JSON file. In the markdown file follow the rules of templating from [https://pkg.go.dev/text/template](pkg.go.dev/text/template). Then run programm with `-d` or `--data` flag set to the path of your file.

## Examples
Edit files in `examples/` folder. Run 

```sh
go run cmd/mdprinter/main.go examples/example.md -s retro -a left -d examples/exampleData.json
``` 

to get the PDF.