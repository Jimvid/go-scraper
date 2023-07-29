package pokemon

import (
	"github.com/gocolly/colly"
	"scraper/utils"
)

// Product represents a struct for Pokemon products.
type Product struct {
	url, image, name, price string
}

// Scraper represents a struct for data helping colly to scrape information
type Scraper struct {
	url           string
	pagesToScrape int8
	products      []Product
}

// Create a new scrape session
func NewPokemonScraper(url string) *Scraper {
	s := &Scraper{
		url:           url,
		pagesToScrape: 1,
	}

	return s
}

func (s *Scraper) PreformScrape(scraper *colly.Collector) {
	scraper.OnHTML("li.product", func(e *colly.HTMLElement) {
		pokemonProduct := Product{
			url:   e.ChildAttr("a", "href"),
			image: e.ChildAttr("img", "src"),
			name:  e.ChildText("h2"),
			price: e.ChildText(".price"),
		}

		s.products = append(s.products, pokemonProduct)
	})

	scraper.Visit(s.url)

	headers := []string{
		"url",
		"image",
		"name",
		"price",
	}

	// utils.WriteToCSV(headers, pokemonProducts, "output/products.csv")
	err := utils.WriteToCSV(headers, convertToCustomData(s.products), "output/pokemonProducts.csv")
	if err != nil {
		panic(err)
	}
}

// If target site is paginated, set limit to pagination
func (s *Scraper) SetPagesToScrape(input int8) {
	s.pagesToScrape = input
}

// ToCSVRecord converts Product to a slice of strings for CSV writing.
func (p Product) ToCSVRecord() []string {
	return []string{
		p.name,
		p.image,
		p.url,
		p.price,
	}
}

// ConvertToCustomData converts a slice of PokemonProduct to a slice of CustomData.
func convertToCustomData(pokemonProducts []Product) []utils.CustomData {
	customDataSlice := make([]utils.CustomData, len(pokemonProducts))
	for i, v := range pokemonProducts {
		customDataSlice[i] = v
	}
	return customDataSlice
}
