package usecase

import (
	"emerson-argueta/m/v2/modules/identity/domain/user"
	"emerson-argueta/m/v2/modules/identity/repository"
	"emerson-argueta/m/v2/shared/infrastructure/http/authorization"
)

// RegisterUsecase performs registering
type RegisterUsecase struct {
	UserRepo             repository.UserRepo
	AuthorizationService authorization.JwtService
}

// NewRegisterUsecase to register user
func NewRegisterUsecase(
	searchRepo repository.UserRepo,
	authorizationService authorization.JwtService,
) *RegisterUsecase {
	return &RegisterUsecase{
		UserRepo:             searchRepo,
		AuthorizationService: authorizationService,
	}
}

// Execute the usecase
func (uc *RegisterUsecase) Execute(dto *RegisterDTO) (apiKey string, err error) {
	email, err := user.NewEmail(&dto.Email)
	if err != nil {
		return "", err
	}
	hashPassword, err := user.NewHashPassword(&dto.Password)
	if err != nil {
		return "", err
	}
	newUser, err := user.New(&user.Fields{Email: &email, HashPassword: &hashPassword}, nil)
	if err != nil {
		return "", err
	}
	newAPIKey, err := uc.AuthorizationService.IssueNewToken(newUser.GetID())
	if err != nil {
		return "nil", err
	}

	if err := uc.UserRepo.CreateUser(newUser); err != nil {
		return "", err
	}

	return newAPIKey, nil
}
