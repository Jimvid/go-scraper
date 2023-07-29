package main

import (
	"github.com/gocolly/colly"
	"scraper/pokemon"
)

func main() {
	// initializing the slice of structs to store the data to scrape
	Scraper := colly.NewCollector()
	Scraper.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36"

	ps := pokemon.NewPokemonScraper("https://scrapeme.live/shop/page/1/")
	ps.PreformScrape(Scraper)
}
