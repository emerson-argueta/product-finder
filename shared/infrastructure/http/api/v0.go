package api

import (
	identity_routes "emerson-argueta/m/v2/modules/identity/infrastructure/http/routes"
	productfinder_routes "emerson-argueta/m/v2/modules/productfinder/infrastructure/http/routes"

	"net/http"
)

// BaseHandler is a collection of all the api handlers.
type BaseHandler struct {
	BasePath string
	*http.ServeMux
}

// NewBaseHandler with basePath
func NewBaseHandler(basePath string) *BaseHandler {
	bh := new(BaseHandler)
	bh.BasePath = basePath

	mux := http.NewServeMux()
	mux.Handle(basePath, bh)

	productFinderHandler := productfinder_routes.NewPrdouctFinderHandler(basePath)
	mux.Handle(basePath+productfinder_routes.ProductFinderURLPrefix+"/", productFinderHandler)

	identityHandler := identity_routes.NewIdentityHandler(basePath)
	mux.Handle(basePath+identity_routes.IdentityURLPrefix+"/", identityHandler)

	bh.ServeMux = mux

	return bh

}
