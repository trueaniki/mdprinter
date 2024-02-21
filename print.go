package mdprinter

import (
	"context"
	_ "embed"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

//go:embed css/modest.css
var modest string

//go:embed css/retro.css
var retro string

//go:embed css/splendor.css
var splendor string

//go:embed css/air.css
var air string

func Print(data []byte, style string) ([]byte, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	// save html to a temp file
	tmpFile, err := os.CreateTemp("", "*.html")
	tmpFile.Write(data)
	if err != nil {
		return nil, err
	}

	s := spinner.New(spinner.CharSets[35], 100*time.Millisecond)
	var buf []byte
	s.Start()
	err = chromedp.Run(ctx,
		chromedp.Navigate("file://"+tmpFile.Name()),
		chromedp.WaitVisible("body"),
		// Add css to the page
		chromedp.Evaluate(`(function() {
			var style = document.createElement('style');
			style.innerHTML = `+"`"+chooseStyle(style)+"`"+`;
			document.head.appendChild(style);
		})()`, nil),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, _, err = page.PrintToPDF().WithPrintBackground(true).Do(ctx)
			if err != nil {
				return err
			}
			return nil
		}))
	s.Stop()
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func chooseStyle(style string) string {
	switch style {
	case "modest":
		return modest
	case "retro":
		return retro
	case "splendor":
		return splendor
	case "air":
		return air
	default:
		return modest
	}
}
