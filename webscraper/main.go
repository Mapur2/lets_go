package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
)

type Stock struct {
	company, price, change string
}

func main() {
	tickers := []string{
		"MSFT",  // Microsoft
		"AAPL",  // Apple
		"GOOGL", // Alphabet
		"AMZN",  // Amazon
		"META",  // Meta
		"NVDA",  // NVIDIA
		"TSLA",  // Tesla
		"AMD",   // AMD
		"INTC",  // Intel
		"CRM",   // Salesforce
	}

	var stocks []Stock

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {

		r.Headers.Set("User-Agent", "Mozilla 5.0")
		fmt.Println("Visiting url ", r.URL)

	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnHTML("section#quote-hdr", func(h *colly.HTMLElement) {
		stock := Stock{}

		stock.company = h.ChildText("h1")
		stock.price = h.ChildText("fin-streamer[data-field='regularMarketPrice']")
		stock.change = h.ChildText("fin-streamer[data-field='regularMarketChangePercent']")
		fmt.Println(stock)
		stocks = append(stocks, stock)
	})

	for _, t := range tickers {
		c.Visit("https://finance.yahoo.com/quote/" + t + "/")
	}
	c.Wait()
	fmt.Println(stocks)
}
