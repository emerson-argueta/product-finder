package controllers

import (
	"emerson-argueta/m/v2/modules/identity/infrastructure/persistence"
	"emerson-argueta/m/v2/shared/infrastructure/http/authorization"
	"log"
)

// NewMock controller is a mock of controllers with mock repo
func NewMock(
	authorizationService authorization.JwtService,
	logger *log.Logger,
) *Controllers {
	controllers := &Controllers{}

	identityRepos := persistence.IdentityReposMock

	controllers.RegisterController = NewRegisterControllerMock(identityRepos, logger, authorizationService)
	controllers.LoginController = NewLoginControllerMock(identityRepos, logger, authorizationService)
	controllers.NewAPIKeyController = NewAPIKeyControllerMock(identityRepos, logger, authorizationService)

	return controllers
}
