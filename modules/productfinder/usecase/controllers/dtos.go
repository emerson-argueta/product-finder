package controllers

import "emerson-argueta/m/v2/modules/productfinder/dtos"

type searchResponse struct {
	Search *dtos.SearchDTO `json:"search,omitempty"`
}
