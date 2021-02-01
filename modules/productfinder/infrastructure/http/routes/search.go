package routes

import "github.com/labstack/echo"

const (
	// SearchURL used for productfinder search
	SearchURL = "/search"
)

func (h *ProductFinderHandler) handleSearch(ctx echo.Context) error {
	return h.Controllers.SearchController.Execute(ctx)
}
