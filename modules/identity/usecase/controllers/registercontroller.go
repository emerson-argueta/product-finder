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

var _ Controller = &registerController{}

type registerController struct {
	Usecase       *usecase.RegisterUsecase
	Logger        *log.Logger
	Authorization authorization.JwtService
}

// NewRegisterController for register usecase
func NewRegisterController(identityRepos *persistence.Services, logger *log.Logger, authorizationService authorization.JwtService) Controller {
	registerUsecase := usecase.NewRegisterUsecase(&identityRepos.User, authorizationService)

	ctrl := &registerController{
		Usecase:       registerUsecase,
		Logger:        logger,
		Authorization: authorizationService,
	}
	return ctrl
}

// Execute the usecase
func (ctrl *registerController) Execute(ctx echo.Context) (e error) {
	var req registerRequest

	// Decode the request.
	if err := ctx.Bind(&req); err != nil || req.Email == nil || req.Password == nil {
		return response.ErrorResponse(ctx.Response().Writer, response.ErrInvalidJSON, http.StatusBadRequest, ctrl.Logger)
	}

	registerDTO := &usecase.RegisterDTO{
		Email:    *req.Email,
		Password: *req.Password,
	}
	switch apiKey, e := ctrl.Usecase.Execute(registerDTO); e {
	case nil:
		response.EncodeJSON(ctx.Response().Writer, &registerResponse{APIKey: &apiKey}, ctrl.Logger)
	case user.ErrUserExists:
		return response.ErrorResponse(ctx.Response().Writer, e, http.StatusConflict, ctrl.Logger)
	case user.ErrUserIncompleteDetails:
		return response.ErrorResponse(ctx.Response().Writer, e, http.StatusBadRequest, ctrl.Logger)
	default:
		return response.ErrorResponse(ctx.Response().Writer, e, http.StatusInternalServerError, ctrl.Logger)
	}

	return nil
}
