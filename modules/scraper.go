package modules

import (
	"context"
	"log"
	"strings"
	"github.com/chromedp/chromedp"
	"os"
	"fmt"
	"encoding/csv"
	"io"
	"strconv"
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

func Csviterate(){
	file, err := os.Open("units.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()


	reader := csv.NewReader(file)

	lineNumber := 0

	records := make ([][]string, 0)
	prices := []float64{}

	for {
		record, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("err while reading CSV file: %s", err)
		}
		lineNumber++

		records = append(records,record)
	}

	for _, v := range records {
		if len(v) != 1 {
			// malformation of file ignore record
			continue
		}
		prices = append(prices, scrape(v[0]))
		
	}
	fmt.Println(sumWithForLoop(prices))
}