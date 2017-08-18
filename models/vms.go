package models

import (
	"net/http"

	"github.com/ECSTeam/pcf-status/helpers"
)

// VMCollectionRoutesDefinition defines the route for the vms
// for info see: https://opsman-dev-api-docs.cfapps.io/?shell#returning-all-vm-types
var VMCollectionRoutesDefinition = helpers.RouteDefinition{
	Method:  http.MethodGet,
	Path:    "/vms",
	APIType: helpers.OpsMan,
	Handler: func(opsman helpers.API) http.HandlerFunc {
		return opsman.CreateHandler(func(req *http.Request, api helpers.API) (interface{}, error) {
			var container struct {
				Vms []struct {
					Name    string `json:"name"`
					RAM     int    `json:"ram"`
					CPU     int    `json:"cpu"`
					Disk    int    `json:"ephemeral_disk"`
					BuildIn bool   `json:"builtin"`
				} `json:"vm_types"`
			}

			err := api.Get("vm_types", &container)

			return container.Vms, err
		})
	},
}
