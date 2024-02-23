package mdprinter

import (
	"context"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func Print(data []byte, css string) ([]byte, error) {
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
			style.innerHTML = `+"`"+css+"`"+`;
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
