package models

import (
	"net/http"

	"github.com/ECSTeam/pcf-status/helpers"
)

// InfoRoutesDefinition defines the route to return info.
var InfoRoutesDefinition = helpers.RouteDefinition{
	Method:  http.MethodGet,
	Path:    "/info",
	APIType: helpers.AppsMan,
	Handler: func(appsman helpers.API) http.HandlerFunc {
		return appsman.CreateHandler(func(req *http.Request, api helpers.API) (data interface{}, err error) {

			// TODO: this is not complete yet.
			err = api.Get("info", &data)
			return data, err
		})
	},
}
