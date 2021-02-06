package routes

import (
	"emerson-argueta/m/v2/modules/identity/usecase/controllers"
	"emerson-argueta/m/v2/shared/infrastructure"
	"emerson-argueta/m/v2/shared/infrastructure/http/authorization"
	"emerson-argueta/m/v2/shared/infrastructure/http/middleware"
	"log"
	"os"

	"github.com/labstack/echo"
)

const (
	// IdentityURLPrefix used for communitygoaltracker routes
	IdentityURLPrefix = ""
)

// IdentityHandler represents an HTTP API handler.
type IdentityHandler struct {
	*echo.Echo
	*controllers.Controllers
	Logger *log.Logger
}

// NewIdentityHandler uses the labstack echo router.
func NewIdentityHandler(apiBaseURL string) *IdentityHandler {
	h := new(IdentityHandler)

	echoRouter := echo.New()
	logger := log.New(os.Stderr, "", log.LstdFlags)

	authorizationService := authorization.NewJWTService(infrastructure.GlobalConfig)
	controllers := controllers.New(authorizationService, logger)

	h.Echo = echoRouter
	h.Logger = logger
	h.Controllers = controllers

	public := h.Group(apiBaseURL + IdentityURLPrefix)
	public.POST(RegisterURL, h.handleRegister)
	public.POST(LoginURL, h.handleLogin)

	restricted := h.Group(apiBaseURL + IdentityURLPrefix)
	restricted.Use(middleware.APIKeyMiddleware)
	restricted.GET(NewAPIKeyURL, h.handleNewAPIKey)

	return h
}
