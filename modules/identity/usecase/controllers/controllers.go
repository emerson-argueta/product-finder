package controllers

import (
	"emerson-argueta/m/v2/modules/identity/infrastructure/persistence"
	"emerson-argueta/m/v2/shared/infrastructure/http/authorization"
	"log"

	"github.com/labstack/echo"
)

// Controller for usescases
type Controller interface {
	Execute(ctx echo.Context) error
}

// Controllers holds all controllers
type Controllers struct {
	RegisterController  Controller
	LoginController     Controller
	NewAPIKeyController Controller
}

// New controller holds all necessary controllers
func New(
	authorizationService authorization.JwtService,
	logger *log.Logger,
) *Controllers {
	controllers := &Controllers{}

	identityRepos := persistence.IdentityRepos

	controllers.RegisterController = NewRegisterController(identityRepos, logger, authorizationService)
	controllers.LoginController = NewLoginController(identityRepos, logger, authorizationService)
	controllers.NewAPIKeyController = NewAPIKeyController(identityRepos, logger, authorizationService)

	return controllers
}
