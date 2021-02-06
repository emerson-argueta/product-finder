package middleware

import (
	"emerson-argueta/m/v2/modules/identity/domain/user"
	"emerson-argueta/m/v2/modules/identity/infrastructure/persistence"
	"emerson-argueta/m/v2/shared/infrastructure"
	"emerson-argueta/m/v2/shared/infrastructure/http/authorization"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// APIKeyMiddleware for api key validation
var APIKeyMiddleware = middleware.KeyAuthWithConfig(
	middleware.KeyAuthConfig{
		KeyLookup: "header:api_key",
		Validator: apiKeyValidator,
	},
)

var apiKeyValidator = func(key string, c echo.Context) (bool, error) {
	jwtService := authorization.NewJWTService(infrastructure.GlobalConfig)

	id, err := jwtService.VerifyTokenAndExtractIDClaim(key)
	if err != nil && err == authorization.ErrAuthorizationFailed {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	_, err = persistence.IdentityRepos.User.RetrieveUserByID(id)
	if err != nil && err == user.ErrUserNotFound {
		return false, nil
	}
	if err != nil {
		return false, err
	}

	return true, nil
}
