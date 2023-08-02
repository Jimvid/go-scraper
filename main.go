package main

import (
	"fmt"
	"scraper/pokemon"

	"github.com/gocolly/colly"
)

func main() {
	// initializing the slice of structs to store the data to scrape
	Scraper := colly.NewCollector()

	// Set user agent
	Scraper.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	// Handle basic errors
	Scraper.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r.StatusCode, "\nError:", err)
	})

	// Scrape pokemon products
	ps := pokemon.NewPokemonScraper("https://scrapeme.live/shop/page/1/")
	ps.SetPageLimit(5)
	ps.PreformScrape(Scraper)
}
