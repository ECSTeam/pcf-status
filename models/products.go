package models

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/ECSTeam/pcf-status/helpers"
)

// OpsManProductCollectionRoute defines the route for the products
// for info see: https://opsman-dev-api-docs.cfapps.io/?shell#deployed-products
var OpsManProductCollectionRoute = helpers.RouteDefinition{
	Method:  http.MethodGet,
	Path:    "/models/products",
	APIType: helpers.OpsMan,
	Handler: func(om helpers.API) http.HandlerFunc {
		return helpers.CreateHandler(om, func(req *http.Request, api helpers.API) (interface{}, error) {

			var err error
			var products helpers.OpsManProducts
			if opsman, ok := api.(*helpers.OpsManAPI); ok {
				if products, err = opsman.GetProducts(); err == nil {
					for index, prod := range products {
						href := helpers.MakeURL(req, "products", prod.GUID)
						products[index].Append("self", href)
					}
				}
			}

			return products, err
		})
	},
}

// OpsManProductRoute defines the process for getting the
// for more info: https://opsman-dev-api-docs.cfapps.io/?shell#retrieving-status-of-product-jobs
var OpsManProductRoute = helpers.RouteDefinition{
	Method:  http.MethodGet,
	Path:    "/models/products/{guid}",
	APIType: helpers.OpsMan,
	Handler: func(opsman helpers.API) http.HandlerFunc {
		return helpers.CreateHandler(opsman, func(req *http.Request, api helpers.API) (interface{}, error) {

			var err error
			var product interface{}

			vars := mux.Vars(req)
			if guid, ok := vars["guid"]; ok {
				path := strings.Join([]string{
					"deployed",
					"products",
					guid,
					"status",
				}, "/")

				err = api.Get(path, &product)
			} else {
				err = helpers.NewError("No id provided.")
			}

			return product, err
		})
	},
}
