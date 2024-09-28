package modules

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"istio.io/pkg/cache"
)

// Create a global cache instance with a TTL of 1 hour
var priceCache = cache.NewTTL(time.Hour, time.Minute)

// sumWithForLoop sums up the elements in a float64 slice.
func sumWithForLoop(numbers []float64) float64 {
	sum := 0.00
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// scrape scrapes the price of the given unit from the Warhammer website.
func scrape(unit string) float64 {

	// Replace spaces in the unit string with dashes to format the URL
	unitbar := strings.ReplaceAll(unit, " ", "-")

	// Check if the price is already in the cache
	if cachedPrice, found := priceCache.Get(unitbar); found {
		// If found in the cache, return the cached price
		return cachedPrice.(float64)
	}

	// Create a new ChromeDP context
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	var price string
	// Navigate and scrape the price
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.warhammer.com/en-US/plp?search="+unitbar),
		chromedp.Poll("document.querySelector('[data-testid=product-cards-container]')?.childElementCount > 0 && parseInt(document.querySelector('[data-testid=results-count]')?.innerHTML) < 3500", nil),
		chromedp.InnerHTML("[data-testid=product-card-current-price]", &price, chromedp.NodeVisible),
		chromedp.ActionFunc(func(ctx context.Context) error {
			return nil
		}),
	)
	if err != nil {
		log.Fatal("Error while performing the automation logic:", err)
	}

	// Convert the price string to a float64, assuming the price string starts with a currency symbol like "$"
	pricefloat, _ := strconv.ParseFloat(price[1:], 64)

	// Store the price in the cache with unitbar as the key
	priceCache.Set(unitbar, pricefloat)

	return pricefloat
}

// Csviterate iterates over the data, scrapes prices, applies the multiplier, and returns the total value.
func Csviterate(data []string, multiplier float64) string {

	// Slice to hold prices
	prices := []float64{}

	// Iterate over the unit names in the data
	for _, v := range data {
		// Append the scraped or cached price to the prices slice
		prices = append(prices, scrape(v))
	}

	// Calculate the total value in dollars
	dollarfloat := sumWithForLoop(prices) * (multiplier * .01)

	// Format the result as a string with two decimal places
	dollarValue := fmt.Sprintf("$%.2f", dollarfloat)

	return dollarValue
}
