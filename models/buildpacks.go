package models

import (
	"net/http"

	"github.com/ECSTeam/pcf-status/helpers"
)

// BuildpacksCollectionRoutesDefinition defines the route to create buildpacks.
var BuildpacksCollectionRoutesDefinition = helpers.RouteDefinition{
	Method:  http.MethodGet,
	Path:    "/buildpacks",
	APIType: helpers.AppsMan,
	Handler: func(appsman helpers.API) http.HandlerFunc {
		return appsman.CreateHandler(func(req *http.Request, api helpers.API) (data interface{}, err error) {
			container := &cfContainer{}
			err = api.Get("buildpacks", container)
			return container.Dump(err), err
		})
	},
}
