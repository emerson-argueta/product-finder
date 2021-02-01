package webscraper_test

import (
	"emerson-argueta/m/v2/modules/productfinder/domain/productfinder"
	"emerson-argueta/m/v2/modules/productfinder/infrastructure/webscraper"
	"testing"
)

func TestWebScraper(t *testing.T) {
	gtin12 := "850008366079"
	inputBarcode, _ := productfinder.NewBarcode(gtin12)
	expectedTitle := "Pure Omega Plus by Nutragen - 120 Softgels"

	searchService := webscraper.ProductFinderRepos
	searchResult, err := searchService.Search.ExecuteSearch(inputBarcode)
	if err != nil {
		t.Fatalf("Unexpected error")
	}
	if len(searchResult.GetProducts()) < 1 {
		t.Fatalf("Did not get expected result")
	}

	outputTitle := searchResult.GetProducts()[0].Title
	if outputTitle != expectedTitle {
		t.Fatalf("Expected %s but got %s", expectedTitle, outputTitle)
	}

}
func TestWebScraperWithNoResults(t *testing.T) {
	gtin12 := "850008366071"
	inputBarcode, _ := productfinder.NewBarcode(gtin12)

	searchService := webscraper.ProductFinderRepos
	searchResult, err := searchService.Search.ExecuteSearch(inputBarcode)
	if err != nil {
		t.Fatalf("Unexpected error")
	}
	if len(searchResult.GetProducts()) != 0 {
		t.Fatalf("Did not get expected result")
	}

}
