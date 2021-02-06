package usecase

import (
	"emerson-argueta/m/v2/modules/identity/domain/user"
	"emerson-argueta/m/v2/modules/identity/repository"
	"emerson-argueta/m/v2/shared/infrastructure/http/authorization"
)

// APIKeyUsecase performs registering
type APIKeyUsecase struct {
	UserRepo             repository.UserRepo
	AuthorizationService authorization.JwtService
}

// NewAPIKeyUsecase to register user
func NewAPIKeyUsecase(
	userRepo repository.UserRepo,
	authorizationService authorization.JwtService,
) *APIKeyUsecase {
	return &APIKeyUsecase{
		UserRepo:             userRepo,
		AuthorizationService: authorizationService,
	}
}

// Execute the usecase
func (uc *APIKeyUsecase) Execute(dto *APIKeyDTO) (apiKey string, err error) {
	toBeFoundEmail, err := user.NewEmail(&dto.Email)
	if err != nil {
		return "", err
	}

	foundUser, err := uc.UserRepo.RetrieveUserByEmail(toBeFoundEmail)
	if err != nil {
		return "", err
	}

	if err = user.CompareHashAndPassword(foundUser.GetHashPassword(), &dto.Password); err != nil {
		return "", err
	}

	return uc.AuthorizationService.IssueNewToken(foundUser.GetID())
}
