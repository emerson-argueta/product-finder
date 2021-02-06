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

var _ Controller = &loginController{}

type loginController struct {
	Usecase       *usecase.LoginUsecase
	Logger        *log.Logger
	Authorization authorization.JwtService
}

// NewLoginController for login usecase
func NewLoginController(identityRepos *persistence.Services, logger *log.Logger, authorizationService authorization.JwtService) Controller {
	loginUsecase := usecase.NewLoginUsecase(&identityRepos.User, authorizationService)

	ctrl := &loginController{
		Usecase:       loginUsecase,
		Logger:        logger,
		Authorization: authorizationService,
	}
	return ctrl
}

// Execute the usecase
func (ctrl *loginController) Execute(ctx echo.Context) (e error) {
	var req loginRequest

	// Decode the request.
	if err := ctx.Bind(&req); err != nil || req.Email == nil || req.Password == nil {
		return response.ErrorResponse(ctx.Response().Writer, response.ErrInvalidJSON, http.StatusBadRequest, ctrl.Logger)
	}

	dto := &usecase.LoginDTO{
		Email:    *req.Email,
		Password: *req.Password,
	}
	switch e := ctrl.Usecase.Execute(dto); e {
	case nil:
		sucessMessage := "login successful"
		response.EncodeJSON(ctx.Response().Writer, &loginResponse{Message: &sucessMessage}, ctrl.Logger)
	case user.ErrUserIncorrectCredentials:
		return response.ErrorResponse(ctx.Response().Writer, e, http.StatusUnauthorized, ctrl.Logger)
	case user.ErrUserNotFound:
		return response.ErrorResponse(ctx.Response().Writer, e, http.StatusNotFound, ctrl.Logger)
	default:
		return response.ErrorResponse(ctx.Response().Writer, e, http.StatusInternalServerError, ctrl.Logger)
	}

	return nil
}
