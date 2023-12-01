package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(colly.AllowedDomains("quotes.toscrape.com"))

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.1.2 Safari/605.1.15")
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Response Code", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("error", err.Error())
	})

	c.OnHTML(".quote", func(h *colly.HTMLElement) {
		div := h.DOM
		text := div.Find(".text").Text()
		author := div.Find(".author").Text()
		fmt.Printf("%s\nBy %s\n\n", text, author)
	})

	c.Visit("http://quotes.toscrape.com/")
}
