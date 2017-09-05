package helpers

import (
	"log"
	"strings"
	"time"
)

// OpsManAPI is the structure for the OpsMan interface.
type OpsManAPI struct {
	baseAPI
	cache Cache
}

// OpsManProducts defines the collection of products.
// Defined by: GET /api/v0/deployed/products
// Documentation: https://opsman-dev-api-docs.cfapps.io/?shell#deployed-products
type OpsManProducts []struct {
	Name string `json:"installation_name"`
	GUID string `json:"guid"`
	Type string `json:"type"`
}

// NewOpsManAPI creates a new api.
func NewOpsManAPI(config OpsManConfig) (api *OpsManAPI, err error) {

	// First, get the UAA, this is done by simply querying the base address.
	var uaaAddress string
	var cache Cache
	if cache, err = NewCache(time.Duration(30 * time.Second)); err == nil {
		if uaaAddress, err = findUAAURL(config.Address); err == nil {
			opsMan := &OpsManAPI{
				baseAPI: baseAPI{
					clientID:   "opsman",
					uaaAddress: uaaAddress,
					address:    config.Address,
					user:       config.User,
					password:   config.Password,
				},
				cache: cache,
			}

			if err = opsMan.bindURLMethod(func(items ...string) (url string) {
				path := strings.Join(items, "/")
				base := opsMan.address + "/api/v0/"

				if !strings.HasPrefix(path, base) {
					url = base + path
				} else {
					url = path
				}

				return url
			}); err == nil {
				api = opsMan
			}
		}
	}

	return api, err
}

// GetCredentials will return the credentials by reference.
func (api *OpsManAPI) GetCredentials(references ...string) (user string, password string, err error) {

	// Type documented here:
	// https://opsman-dev-api-docs.cfapps.io/?shell#fetching-credentials
	var credentials struct {
		Creds struct {
			Value struct {
				Identity string `json:"identity"`
				Password string `json:"password"`
			} `json:"value"`
		} `json:"credential"`
	}

	if err = api.Get(strings.Join(references, "/"), &credentials); err == nil {
		user = credentials.Creds.Value.Identity
		password = credentials.Creds.Value.Password
	}
	return user, password, err
}

// GetDiagnosticReport returns the collection of products.
func (api *OpsManAPI) GetDiagnosticReport() (report DiagnosticReport, err error) {

	label := "diagnostic"
	if raw, has := api.cache.Get(label); !has {
		if err = api.Get("diagnostic_report", &report); err == nil {
			api.cache.Add(report, label)
		}
	} else {
		var ok bool
		if report, ok = raw.(DiagnosticReport); !ok {
			err = Error("Unacceptable item from cache.")
		}
	}
	return report, err
}

// GetProducts returns the collection of products.
func (api *OpsManAPI) GetProducts() (products OpsManProducts, err error) {
	err = api.Get("deployed/products", &products)
	return products, err
}

// GetAllManifests will return all the manifests.
func (api *OpsManAPI) GetAllManifests() (manifests []Manifest, err error) {

	var products OpsManProducts
	if products, err = api.GetProducts(); err == nil {

		for _, prod := range products {
			var manifest Manifest
			if manifest, err = api.GetProductManifest(prod.Type); err == nil {
				manifests = append(manifests, manifest)
			} else {
				break
			}
		}
	}

	return manifests, err
}

// GetProductManifest returns the product manifest.
func (api *OpsManAPI) GetProductManifest(product string) (manifest Manifest, err error) {

	// First clean out the cached items that are stale.
	api.cache.Clear()

	label := strings.Join([]string{"manifest", product}, "/")
	if cached, has := api.cache.Get(label); !has {
		log.Printf("Getting manifest: %s", product)
		var products OpsManProducts
		if products, err = api.GetProducts(); err == nil {
			for _, prod := range products {
				if prod.Type == product {
					if err = api.Get(api.MakeAPIURL("deployed", "products", prod.GUID, "manifest"), &manifest); err == nil {
						api.cache.Add(manifest, label)
					}

					break
				}
			}
		}
	} else {
		var ok bool
		if manifest, ok = cached.(Manifest); !ok {
			err = Error("Unacceptable item from cache.")
		}
	}

	return manifest, err
}
