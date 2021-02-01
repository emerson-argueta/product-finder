package webscraper

import (
	"bytes"
	"emerson-argueta/m/v2/modules/productfinder/domain/productfinder"
	"emerson-argueta/m/v2/modules/productfinder/repository"
	"encoding/json"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// SearchType for google to be matched
type SearchType string

func (s *SearchType) String() string {
	return string(*s)
}

const (
	// Shopping search engine
	Shopping = SearchType("shop")
	// SearchBaseURL for google search
	SearchBaseURL = "https://www.google.com/search"
	// APIBaseURL for web-scraper api
	APIBaseURL = "http://localhost:5000/api/v1/search"
	// UserAgent for scraping
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36"
)

var _ repository.SearchRepo = &Search{}

// Search holds reference to database client.
type Search struct {
	Client *http.Request
}

// ExecuteSearch to find products
func (s *Search) ExecuteSearch(barcode *productfinder.Barcode) (productfinder.Search, error) {
	params := map[string]string{
		"q":   barcode.Gtin12,
		"tbm": string(Shopping),
	}
	requestBody, err := json.Marshal(map[string]interface{}{"baseURL": SearchBaseURL, "queryParams": params})
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("POST", APIBaseURL, bytes.NewBuffer(requestBody))
	req.Header.Set("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err

	}
	products, err := parseWebScrapeResults(doc)
	if err != nil {
		return nil, err
	}
	searchFields := &productfinder.Fields{
		Barcode:  barcode,
		Products: products,
	}

	return productfinder.CreateSearch(searchFields)
}
