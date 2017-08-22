package models

import (
	"net/http"

	"github.com/ECSTeam/pcf-status/helpers"
)

// OpsManVMTypesRoute defines the route for the vms_types
// for info see: https://opsman-dev-api-docs.cfapps.io/?shell#returning-all-vm-types
var OpsManVMTypesRoute = helpers.RouteDefinition{
	Method:  http.MethodGet,
	Path:    "/models/vm_types",
	APIType: helpers.OpsMan,
	Handler: func(opsman helpers.API) http.HandlerFunc {
		return helpers.CreateHandler(opsman, func(req *http.Request, api helpers.API) (interface{}, error) {
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

type vmRecord struct {
	Product string `json:"product"`

	// stemcell, vm-type, count
	Instances map[string]map[string]int `json:"instances"`
}

// OpsManVMInstances will return all the vm instances.
var OpsManVMInstances = helpers.RouteDefinition{
	Method:  http.MethodGet,
	Path:    "/models/vms",
	APIType: helpers.OpsMan,
	Handler: func(opsman helpers.API) http.HandlerFunc {
		return helpers.CreateHandler(opsman, func(req *http.Request, api helpers.API) (data interface{}, err error) {
			if opsman, ok := api.(*helpers.OpsManAPI); ok {
				var manifests []helpers.Manifest

				var record []vmRecord
				if manifests, err = opsman.GetAllManifests(); err == nil {

					for _, manifest := range manifests {

						entry := vmRecord{
							Product:   manifest.Name,
							Instances: map[string]map[string]int{},
						}

						for _, group := range manifest.InstanceGroup {
							if _, has := entry.Instances[group.Stemcell]; !has {
								entry.Instances[group.Stemcell] = map[string]int{}
							}

							if _, has := entry.Instances[group.Stemcell][group.VMType]; !has {
								entry.Instances[group.Stemcell][group.VMType] = group.Instances
							} else {
								entry.Instances[group.Stemcell][group.VMType] += group.Instances
							}
						}

						if len(entry.Instances) > 0 {
							record = append(record, entry)
						}
					}
				}

				if err == nil {
					data = record
				}
			}

			return data, err
		})
	},
}
