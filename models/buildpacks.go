package models

import (
	"net/http"

	"github.com/ECSTeam/pcf-status/helpers"
)

// AppsManBuildpacksRoute defines the route to create buildpacks.
var AppsManBuildpacksRoute = helpers.RouteDefinition{
	Method:  http.MethodGet,
	Path:    "/models/buildpacks",
	APIType: helpers.AppsMan,
	Handler: func(appsman helpers.API) http.HandlerFunc {
		return helpers.CreateHandler(appsman, func(req *http.Request, api helpers.API) (data interface{}, err error) {
			container := &cfContainer{}
			err = api.Get("buildpacks", container)
			return container.Dump(err), err
		})
	},
}
