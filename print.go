package mdprinter

import (
	"context"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func Print(data []byte) ([]byte, error) {
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		chromedp.Flag("disable-gpu", true),
		chromedp.WindowSize(1920, 1080),
	)...)
	defer cancel()
	ctx, cancel := chromedp.NewContext(allocCtx)
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
		chromedp.Sleep(1*time.Second),
		chromedp.ActionFunc(func(ctx context.Context) error {
			printAction := page.
				PrintToPDF()
			buf, _, err = printAction.Do(ctx)
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
