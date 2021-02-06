package usecase

import (
	"emerson-argueta/m/v2/modules/productfinder/domain/productfinder"
	"emerson-argueta/m/v2/modules/productfinder/repository"
	"emerson-argueta/m/v2/shared/infrastructure/http/authorization"
)

// SearchUsecase performs registering
type SearchUsecase struct {
	SearchRepo           repository.SearchRepo
	AuthorizationService authorization.JwtService
}

// NewSearchUsecase to register user
func NewSearchUsecase(
	searchRepo repository.SearchRepo,
	authorizationService authorization.JwtService,
) *SearchUsecase {
	return &SearchUsecase{
		SearchRepo:           searchRepo,
		AuthorizationService: authorizationService,
	}
}

// Execute the usecase
func (uc *SearchUsecase) Execute(dto *SearchDTO) (productfinder.Search, error) {
	barcode, err := productfinder.NewBarcode(dto.Barcode)
	if err != nil {
		return nil, err
	}

	return uc.SearchRepo.ExecuteSearch(barcode)
}
