package models

import (
	"net/http"

	"github.com/ECSTeam/pcf-status/helpers"
)

// AppsManStemcellsRoute defines the route fo the stemcells
var AppsManStemcellsRoute = helpers.RouteDefinition{
	Method:  http.MethodGet,
	Path:    "/models/stemcells",
	APIType: helpers.OpsMan,
	Handler: func(om helpers.API) http.HandlerFunc {
		return helpers.CreateHandler(om, func(req *http.Request, api helpers.API) (data interface{}, err error) {
			if opsman, ok := api.(*helpers.OpsManAPI); ok {
				var manifests []helpers.Manifest
				if manifests, err = opsman.GetAllManifests(); err == nil {

					stemcells := []helpers.Stemcell{}
					for _, manifest := range manifests {

						// cleanup the data.
						for i := range manifest.Stemcells {
							manifest.Stemcells[i].Source = manifest.Name
						}

						stemcells = append(stemcells, manifest.Stemcells...)
					}

					data = stemcells
				}
			} else {
				err = helpers.NewError("Invalid API.")
			}

			return data, err
		})
	},
}
