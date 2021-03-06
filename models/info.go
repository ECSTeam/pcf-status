package models

import (
	"net/http"

	"github.com/ECSTeam/pcf-status/helpers"
)

// FoundationInfo defines the information
type FoundationInfo struct {
	Type    string `json:"iaas-type"`
	Version string `json:"version"`
}

// AppsManInfoRoute defines the route to return info.
var AppsManInfoRoute = helpers.RouteDefinition{
	Method:  http.MethodGet,
	Path:    "/models/info",
	APIType: helpers.OpsMan,
	Handler: func(appsman helpers.API) http.HandlerFunc {
		return helpers.CreateHandler(appsman, func(req *http.Request, api helpers.API) (data interface{}, err error) {
			if opsman, ok := api.(*helpers.OpsManAPI); ok {
				var report helpers.DiagnosticReport
				if report, err = opsman.GetDiagnosticReport(); err == nil {
					data = FoundationInfo{
						Type:    report.Type,
						Version: report.Versions.Release,
					}
				}
			} else {
				err = helpers.NewError("Invalid API.")
			}

			return data, err
		})
	},
}
