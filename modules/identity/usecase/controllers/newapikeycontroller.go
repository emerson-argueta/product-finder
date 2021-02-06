package controllers

import (
	"emerson-argueta/m/v2/modules/identity/domain/user"
	"emerson-argueta/m/v2/modules/identity/infrastructure/persistence"
	"emerson-argueta/m/v2/modules/identity/usecase"
	"emerson-argueta/m/v2/shared/infrastructure/http/authorization"
	"emerson-argueta/m/v2/shared/infrastructure/http/response"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

var _ Controller = &apiKeyController{}

type apiKeyController struct {
	Usecase       *usecase.APIKeyUsecase
	Logger        *log.Logger
	Authorization authorization.JwtService
}

// NewAPIKeyController for search usecase
func NewAPIKeyController(identityRepos *persistence.Services, logger *log.Logger, authorizationService authorization.JwtService) Controller {
	apiKeyUsecase := usecase.NewAPIKeyUsecase(&identityRepos.User, authorizationService)

	ctrl := &apiKeyController{
		Usecase:       apiKeyUsecase,
		Logger:        logger,
		Authorization: authorizationService,
	}
	return ctrl
}

// Execute the usecase
func (ctrl *apiKeyController) Execute(ctx echo.Context) error {
	var req apiKeyRequest

	// Decode the request.
	if err := ctx.Bind(&req); err != nil || req.Email == nil || req.Password == nil {
		return response.ErrorResponse(ctx.Response().Writer, response.ErrInvalidJSON, http.StatusBadRequest, ctrl.Logger)
	}

	dto := &usecase.APIKeyDTO{
		Email:    *req.Email,
		Password: *req.Password,
	}
	switch newAPIKey, e := ctrl.Usecase.Execute(dto); e {
	case nil:
		response.EncodeJSON(ctx.Response().Writer, &apiKeyResponse{APIKey: &newAPIKey}, ctrl.Logger)
	case user.ErrUserIncorrectCredentials:
		return response.ErrorResponse(ctx.Response().Writer, e, http.StatusUnauthorized, ctrl.Logger)
	case user.ErrUserNotFound:
		return response.ErrorResponse(ctx.Response().Writer, e, http.StatusNotFound, ctrl.Logger)
	default:
		return response.ErrorResponse(ctx.Response().Writer, e, http.StatusInternalServerError, ctrl.Logger)
	}

	return nil
}
