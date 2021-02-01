package webscraper

import (
	"emerson-argueta/m/v2/modules/productfinder/domain/productfinder"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// ProductFinderRepos implementation
var ProductFinderRepos = new()

// Service for product finder
type Service struct {
	Search Search
}

func new() *Service {
	service := &Service{}

	req, _ := http.NewRequest("POST", APIBaseURL, nil)

	service.Search.Client = req

	return service
}
func parseWebScrapeResults(doc *goquery.Document) ([]*productfinder.Product, error) {

	shoppingItems := doc.Find("a.translate-content")

	gg := []*productfinder.Product{}
	shoppingItems.Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		g := &productfinder.Product{Title: title}
		gg = append(gg, g)

	})

	return gg, nil
}
