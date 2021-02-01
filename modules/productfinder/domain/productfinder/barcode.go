package productfinder

// Barcode for products
type Barcode struct {
	Gtin12 string
}

// NewBarcode with gtin12 number
func NewBarcode(gtin12 string) (*Barcode, error) {

	if err := validateGtin12(gtin12); err != nil {
		return nil, err
	}
	barcode := &Barcode{Gtin12: gtin12}

	return barcode, nil
}

func validateGtin12(gtin12 string) error {
	return nil
}
