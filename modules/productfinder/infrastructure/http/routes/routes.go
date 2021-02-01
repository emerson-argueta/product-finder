package routes

import (
	"emerson-argueta/m/v2/modules/productfinder/usecase/controllers"
	"emerson-argueta/m/v2/shared/infrastructure/http/authorization"
	"log"
	"os"

	"github.com/labstack/echo"
)

const (
	// ProductFinderURLPrefix used for communitygoaltracker routes
	ProductFinderURLPrefix = "/productfinder"
)

// ProductFinderHandler represents an HTTP API handler.
type ProductFinderHandler struct {
	*echo.Echo
	*controllers.Controllers
	Logger *log.Logger
}

// NewPrdouctFinderHandler uses the labstack echo router.
func NewPrdouctFinderHandler(apiBaseURL string) *ProductFinderHandler {
	h := new(ProductFinderHandler)

	echoRouter := echo.New()
	logger := log.New(os.Stderr, "", log.LstdFlags)

	authorizationService := authorization.AuthorizationService
	controllers := controllers.New(authorizationService, logger)

	h.Echo = echoRouter
	h.Logger = logger
	h.Controllers = controllers

	restricted := h.Group(apiBaseURL + ProductFinderURLPrefix)
	// restricted.Use(middleware.JwtMiddleware)
	restricted.GET(SearchURL, h.handleSearch)

	return h
}
