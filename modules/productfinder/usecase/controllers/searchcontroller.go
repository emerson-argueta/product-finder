package controllers

import (
	"emerson-argueta/m/v2/modules/productfinder/domain/productfinder"
	"emerson-argueta/m/v2/modules/productfinder/dtos"
	"emerson-argueta/m/v2/modules/productfinder/infrastructure/webscraper"
	"emerson-argueta/m/v2/modules/productfinder/usecase"
	"emerson-argueta/m/v2/shared/infrastructure/http/authorization"
	"emerson-argueta/m/v2/shared/infrastructure/http/response"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

var _ Controller = &searchController{}

type searchController struct {
	Usecase       *usecase.SearchUsecase
	Logger        *log.Logger
	Authorization authorization.JwtService
}

// NewSearchController for search usecase
func NewSearchController(productFinderRepo *webscraper.Service, logger *log.Logger, authorizationService authorization.JwtService) Controller {
	searchUsecase := usecase.NewSearchUsecase(&productFinderRepo.Search, authorizationService)

	ctrl := &searchController{
		Usecase:       searchUsecase,
		Logger:        logger,
		Authorization: authorizationService,
	}
	return ctrl
}

// Execute the usecase
func (ctrl *searchController) Execute(ctx echo.Context) (e error) {
	barcode := ctx.QueryParam("barcode")
	if barcode == "" {
		return response.ErrorResponse(ctx.Response().Writer, productfinder.ErrSearchIncompleteDetails, http.StatusBadRequest, ctrl.Logger)
	}

	// extract user id from authKey stored by JwtMiddleware handler func
	// authKey := ctx.Get("user")
	// userID, e := ctrl.Authorization.JwtService().Authorize(authKey)
	// if e != nil || userID == nil {
	// 	return response.ErrorResponse(ctx.Response().Writer, e, http.StatusInternalServerError, ctrl.Logger)
	// }

	dto := &usecase.SearchDTO{Barcode: barcode}
	switch searchResult, err := ctrl.Usecase.Execute(dto); err {
	case nil:
		searchDTO := dtos.SearchToDTO(searchResult)
		response.EncodeJSON(ctx.Response().Writer, &searchResponse{Search: searchDTO}, ctrl.Logger)
	default:
		return response.ErrorResponse(ctx.Response().Writer, e, http.StatusInternalServerError, ctrl.Logger)
	}

	return nil
}
