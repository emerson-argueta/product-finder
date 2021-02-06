package routes

import "github.com/labstack/echo"

const (
	// RegisterURL used for register
	RegisterURL = "/register"
	// LoginURL used for login
	LoginURL = "/login"
	// NewAPIKeyURL used for acquiring a new api key
	NewAPIKeyURL = "/newkey"
)

func (h *IdentityHandler) handleRegister(ctx echo.Context) error {
	return h.Controllers.RegisterController.Execute(ctx)
}
func (h *IdentityHandler) handleLogin(ctx echo.Context) error {
	return h.Controllers.LoginController.Execute(ctx)
}
func (h *IdentityHandler) handleNewAPIKey(ctx echo.Context) error {
	return h.Controllers.NewAPIKeyController.Execute(ctx)
}
