package models

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/ECSTeam/pcf-status/helpers"
)

/*
/////////
	"/api/v0/deployed/products/:product-guid/variables"

	{
	  "variables": ["first-variable", "second-variable", "third-variable"]
	}

/////////
	"/api/v0/deployed/products/:product-guid/variables?name=:variable_name"

	{
  	"credhub-password": "example-password"
	}

/////////
	"/api/v0/deployed/products/:product_guid/status"

	{
	  "status": [
	    {
	      "job-name": "web_server-7f841fc2af9c2b357cc4",
	      "index": 0,
	      "az_guid": "ee61aa1e420ed3fdf276",
	      "az_name": "first-az",
	      "ips": [
	        "10.85.42.58"
	      ],
	      "cid": "vm-448ef313-86ee-4049-87cf-764ca2fa97e7",
	      "load_avg": [
	        "0.00",
	        "0.01",
	        "0.03"
	      ],
	      "cpu": {
	        "sys": "0.1",
	        "user": "0.2",
	        "wait": "0.3"
	      },
	      "memory": {
	        "kb": "60632",
	        "percent": "6"
	      },
	      "swap": {
	        "kb": "0",
	        "percent": "0"
	      },
	      "system_disk": {
	        "inode_percent": "31",
	        "percent": "42"
	      },
	      "ephemeral_disk": {
	        "inode_percent": "0",
	        "percent": "1"
	      },
	      "persistent_disk": {
	        "inode_percent": "0",
	        "percent": "0"
	      }
	    }
	  ]
	}

/////////
"/api/v0/deployed/products/:product_guid/static_ips"
[
{
	"name": "job-type1-guid-partition-default-az-guid",
	"ips": [
		"192.168.163.4"
	]
},
{
	"name": "credentials-job-guid-partition-default-az-guid",
	"ips": [
		"192.168.163.7"
	]
}
]

//////////
/api/v0/disk_types

{
  "disk_types": [
    {
      "name": "1024",
      "builtin": true,
      "size_mb": 1024
    },
    {
      "name": "2048",
      "builtin": true,
      "size_mb": 2048
    },
    {
      "name": "5120",
      "builtin": true,
      "size_mb": 5120
    }
  ]
}



//////////////
/api/v0/vm_types

{
  "vm_types": [
    {
      "name": "nano",
      "ram": 512,
      "cpu": 1,
      "ephemeral_disk": 1024,
      "builtin": true
    },
    {
      "name": "micro",
      "ram": 1024,
      "cpu": 1,
      "ephemeral_disk": 2048,
      "builtin": true
    },
    {
      "name": "small.disk",
      "ram": 2048,
      "cpu": 1,
      "ephemeral_disk": 16384,
      "builtin": true
      }
  ]
}


///////////////////
/api/v0/deployed/products/:product_guid/manifest

/api/v0/diagnostic_report
*/

// ProductCollectionRoutesDefinition defines the route for the products
// for info see: https://opsman-dev-api-docs.cfapps.io/?shell#deployed-products
var ProductCollectionRoutesDefinition = helpers.RouteDefinition{
	Method:  http.MethodGet,
	Path:    "/products",
	APIType: helpers.OpsMan,
	Handler: func(opsman helpers.API) http.HandlerFunc {
		return opsman.CreateHandler(func(req *http.Request, api helpers.API) (interface{}, error) {
			var products []struct {
				helpers.Links
				Name string `json:"installation_name"`
				GUID string `json:"guid"`
				Type string `json:"type"`
			}

			err := api.Get("deployed/products", &products)
			for index, prod := range products {
				href := helpers.MakeURL(req, "products", prod.GUID)
				products[index].Append("self", href)
			}

			return products, err
		})
	},
}

// ProductRoutesDefinition defines the process for getting the
var ProductRoutesDefinition = helpers.RouteDefinition{
	Method:  http.MethodGet,
	Path:    "/products/{guid}",
	APIType: helpers.OpsMan,
	Handler: func(opsman helpers.API) http.HandlerFunc {
		return opsman.CreateHandler(func(req *http.Request, api helpers.API) (interface{}, error) {

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

// https://opsman-dev-api-docs.cfapps.io/?shell#retrieving-status-of-product-jobs

// https://opsman-dev-api-docs.cfapps.io/?shell#listing-static-ip-assignments-for-product-jobs

// https://opsman-dev-api-docs.cfapps.io/?shell#retrieving-manifest-for-a-deployed-product
// GET /api/v0/deployed/products/:product_guid/manifest
