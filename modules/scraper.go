package modules

import (
	"context"
	"log"
	"strings"
	"github.com/chromedp/chromedp"
	"strconv"
	"fmt"
)
func sumWithForLoop(numbers []float64) float64 {
    sum := 0.00
    for _, num := range numbers {
        sum += num
    }
    return sum
}

func scrape(unit string) float64 {
	// initialize a controllable Chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	// to release the browser resources when
	// it is no longer needed
	defer cancel()
	var price string

	err := chromedp.Run(ctx,
		// visit the target page
		chromedp.Navigate("https://www.warhammer.com/en-US/plp?search=" + strings.ReplaceAll(unit, " ", "-")),
		// wait for the page to load
		chromedp.Poll("document.querySelector('[data-testid=product-cards-container]')?.childElementCount > 0 && parseInt(document.querySelector('[data-testid=results-count]')?.innerHTML) < 3500", nil),
		chromedp.InnerHTML("[data-testid=product-card-current-price]", &price, chromedp.NodeVisible),
		chromedp.ActionFunc(func(ctx context.Context) error {
			return nil
		}),
	)
	if err != nil {
		log.Fatal("Error while performing the automation logic:", err)
	}
	pricefloat, _ := strconv.ParseFloat(price[1:], 64)
	return pricefloat
}

func Csviterate(data []string) string{

	prices := []float64{}


	for _, v := range data {
		prices = append(prices, scrape(v))
		
	}
	dollarValue := fmt.Sprintf("$%.2f", sumWithForLoop(prices))

	return dollarValue

}