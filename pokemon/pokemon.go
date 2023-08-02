package pokemon

import (
	"github.com/gocolly/colly"
	"scraper/utils"
)

type Product struct {
	url, image, name, price string
}

type Scraper struct {
	url              string
	pagesToScrape    []string
	pagesDiscovered  []string
	pageLimit        int8
	products         []Product
	currentIteration int8
	currentPage      string
}

func NewPokemonScraper(url string) *Scraper {
	s := &Scraper{
		url:              url,
		currentPage:      url,
		pageLimit:        1,
		currentIteration: 1,
		pagesToScrape:    []string{},
		pagesDiscovered:  []string{},
	}

	return s
}

func (s *Scraper) PreformScrape(scraper *colly.Collector) {
	scraper.OnHTML("a.page-numbers", func(e *colly.HTMLElement) {
		newPaginationLink := e.Attr("href")

		if !containsStringInSlice(s.pagesToScrape, newPaginationLink) {
			if !containsStringInSlice(s.pagesDiscovered, newPaginationLink) {
				s.pagesToScrape = append(s.pagesToScrape, newPaginationLink)
			}
			s.pagesDiscovered = append(s.pagesDiscovered, newPaginationLink)
		}
	})

	scraper.OnHTML("li.product", func(e *colly.HTMLElement) {
		pokemonProduct := Product{
			url:   e.ChildAttr("a", "href"),
			image: e.ChildAttr("img", "src"),
			name:  e.ChildText("h2"),
			price: e.ChildText(".price"),
		}

		s.products = append(s.products, pokemonProduct)
	})

	scraper.OnScraped(func(response *colly.Response) {
		// until there is still a page to scrape
		if len(s.pagesToScrape) != 0 && s.currentIteration < s.pageLimit {
			// getting the current page to scrape and removing it from the list
			s.currentPage = s.pagesToScrape[0]
			s.pagesToScrape = s.pagesToScrape[1:]

			// incrementing the iteration counter
			s.currentIteration++

			// visiting a new page
			scraper.Visit(s.currentPage)
		}
	})

	scraper.Visit(s.url)

	headers := []string{
		"url",
		"image",
		"name",
		"price",
	}

	err := utils.WriteToCSV(headers, convertToCustomData(s.products), "output/pokemonProducts.csv")
	if err != nil {
		panic(err)
	}
}

func (s *Scraper) SetPageLimit(input int8) {
	s.pageLimit = input
}

func (p Product) ToCSVRecord() []string {
	return []string{
		p.name,
		p.image,
		p.url,
		p.price,
	}
}

func convertToCustomData(pokemonProducts []Product) []utils.CustomData {
	customDataSlice := make([]utils.CustomData, len(pokemonProducts))
	for i, v := range pokemonProducts {
		customDataSlice[i] = v
	}
	return customDataSlice
}

func containsStringInSlice(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
