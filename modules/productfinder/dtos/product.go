package dtos

import "emerson-argueta/m/v2/modules/productfinder/domain/productfinder"

// ProductDTO for data transer
type ProductDTO struct {
	Title string `json:"title,omitempty"`
}

// ProductToDTO from domain barcode
func ProductToDTO(product *productfinder.Product) *ProductDTO {
	productDTO := &ProductDTO{Title: product.Title}
	return productDTO
}
