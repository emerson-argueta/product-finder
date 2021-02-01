package dtos

import "emerson-argueta/m/v2/modules/productfinder/domain/productfinder"

// SearchDTO for data transfer
type SearchDTO struct {
	Barcode  *BarcodeDTO   `json:"barcode,omitempty"`
	Products []*ProductDTO `json:"products,omitempty"`
}

// SearchToDTO from domain barcode
func SearchToDTO(search productfinder.Search) *SearchDTO {
	barcode := BarcodeToDTO(search.GetBarcode())
	products := search.GetProducts()

	productDTOs := make([]*ProductDTO, len(products))
	for i, product := range products {
		productDTOs[i] = ProductToDTO(product)
	}

	searchDTO := &SearchDTO{
		Barcode:  barcode,
		Products: productDTOs,
	}
	return searchDTO
}
