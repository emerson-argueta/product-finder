package productfinder

// Fields for search
type Fields struct {
	Barcode  *Barcode
	Products []*Product
}

type search struct {
	Fields
}

// Search for products by barcode
type Search interface {
	GetBarcode() *Barcode
	GetProducts() []*Product
}

// CreateSearch for products
func CreateSearch(fields *Fields) (Search, error) {
	search := &search{}
	if fields.Barcode == nil || fields.Products == nil {
		return nil, ErrSearchIncompleteDetails
	}
	search.Barcode = fields.Barcode
	search.Products = fields.Products

	return search, nil
}

func (s *search) GetBarcode() *Barcode {
	return s.Barcode
}
func (s *search) GetProducts() []*Product {
	return s.Products
}
