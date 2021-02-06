package usecase

import (
	"emerson-argueta/m/v2/modules/identity/domain/user"
	"emerson-argueta/m/v2/modules/identity/repository"
	"emerson-argueta/m/v2/shared/infrastructure/http/authorization"
)

// LoginUsecase performs registering
type LoginUsecase struct {
	UserRepo             repository.UserRepo
	AuthorizationService authorization.JwtService
}

// NewLoginUsecase to register user
func NewLoginUsecase(
	UserRepo repository.UserRepo,
	authorizationService authorization.JwtService,
) *LoginUsecase {
	return &LoginUsecase{
		UserRepo:             UserRepo,
		AuthorizationService: authorizationService,
	}
}

// Execute the usecase
func (uc *LoginUsecase) Execute(dto *LoginDTO) error {
	toBeFoundEmail, err := user.NewEmail(&dto.Email)
	if err != nil {
		return err
	}

	foundUser, err := uc.UserRepo.RetrieveUserByEmail(toBeFoundEmail)
	if err != nil {
		return err
	}

	return user.CompareHashAndPassword(foundUser.GetHashPassword(), &dto.Password)
}
