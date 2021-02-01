package repository

import "emerson-argueta/m/v2/modules/productfinder/domain/productfinder"

// SearchRepo used to excute searches for products.
type SearchRepo interface {
	ExecuteSearch(barcode *productfinder.Barcode) (productfinder.Search, error)
}
