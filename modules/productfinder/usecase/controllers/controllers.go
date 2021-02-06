package controllers

import (
	"emerson-argueta/m/v2/modules/productfinder/infrastructure/webscraper"
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
	SearchController Controller
}

// New controller holds all necessary controllers
func New(
	authorizationService authorization.JwtService,
	logger *log.Logger,
) *Controllers {
	controllers := &Controllers{}

	productfinderRepo := webscraper.ProductFinderRepos

	controllers.SearchController = NewSearchController(productfinderRepo, logger, authorizationService)

	return controllers
}
