package modules

import (
	"context"
	"log"
	"fmt"
	"github.com/chromedp/cdproto/dom"
    "github.com/chromedp/chromedp"
	"time"
)

func Scrape() {
	// initialize a controllable Chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	 )
	 // to release the browser resources when
	 // it is no longer needed
	 defer cancel()
 
	 var html string
	 err := chromedp.Run(ctx,
		// visit the target page
		chromedp.Navigate("https://www.warhammer.com/en-US/plp?search=lord-kroak"),
		// wait for the page to load
		chromedp.Sleep(2000*time.Millisecond),
		// extract the raw HTML from the page
		chromedp.ActionFunc(func(ctx context.Context) error {
		   // select the root node on the page
		   rootNode, err := dom.GetDocument().Do(ctx)
		   if err != nil {
			  return err
		   }
		   html, err = dom.GetOuterHTML().WithNodeID(rootNode.NodeID).Do(ctx)
		   return err
		}),
	 )
	 if err != nil {
		log.Fatal("Error while performing the automation logic:", err)
	 }
 
	 fmt.Println(html)
 }

