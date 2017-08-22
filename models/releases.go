package models

import (
	"net/http"

	"github.com/ECSTeam/pcf-status/helpers"
)

// AppsManReleasesRoute defines the route to create releases.
var AppsManReleasesRoute = helpers.RouteDefinition{
	Method:  http.MethodGet,
	Path:    "/models/releases",
	APIType: helpers.OpsMan,
	Handler: func(appsman helpers.API) http.HandlerFunc {
		return helpers.CreateHandler(appsman, func(req *http.Request, api helpers.API) (data interface{}, err error) {
			if opsman, ok := api.(*helpers.OpsManAPI); ok {
				var report helpers.DiagnosticReport
				if report, err = opsman.GetDiagnosticReport(); err == nil {
					data = report.Releases
				}
			} else {
				err = helpers.NewError("Invalid API.")
			}

			return data, err
		})
	},
}
