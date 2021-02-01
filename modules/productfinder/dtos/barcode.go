package dtos

import "emerson-argueta/m/v2/modules/productfinder/domain/productfinder"

// BarcodeDTO for data transfer
type BarcodeDTO struct {
	Gtin12 string `json:"gtin12,omitempty"`
}

// BarcodeToDTO from domain barcode
func BarcodeToDTO(barcode *productfinder.Barcode) *BarcodeDTO {
	barcodeDTO := &BarcodeDTO{barcode.Gtin12}
	return barcodeDTO
}
