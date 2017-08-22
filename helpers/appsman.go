package helpers

import (
	"log"
	"strings"
)

const (

	// CFProductName defines the product namd for cloud foundry.
	CFProductName = "cf"
)

// appsManAPI is the structure for the AppsMan interface.
type appsManAPI struct {
	baseAPI
}

// NewAppsManAPI creates a new api.
func NewAppsManAPI(opsMan *OpsManAPI) (api API, err error) {

	if opsMan != nil {
		var manifest Manifest
		if manifest, err = opsMan.GetProductManifest(CFProductName); err == nil {
			var dnsAddr string
			for _, item := range manifest.InstanceGroup {
				if item.Name == "cloud_controller" {
					if routeReg, ok := item.Properties["route_registrar"]; ok {
						if routeRegMapping, ok := routeReg.(map[string]interface{}); ok {
							if routes, ok := routeRegMapping["routes"]; ok {
								if routeFields, ok := routes.([]interface{}); ok {
									for _, fieldStruct := range routeFields {
										if fields, ok := fieldStruct.(map[string]interface{}); ok {
											if name, has := fields["name"]; has && name == "api" {
												if rawAddresses, has := fields["uris"]; has {
													if addresses, ok := rawAddresses.([]interface{}); ok {
														first := addresses[0]
														if addr, ok := first.(string); ok {
															dnsAddr = addr
															log.Printf("Found the Apps Man root address: %v", dnsAddr)
														}
													}
												}
											}

											if len(dnsAddr) > 0 {
												break
											}
										}
									}
								}
							}
						}
					}

					if len(dnsAddr) > 0 {
						break
					}
				}
			}

			// The dnsAddr is just the DNS name.
			var address string
			if !strings.HasPrefix(dnsAddr, "http") {
				testedAddr := "https://" + dnsAddr

				// The address is acceptable.
				err = callURL("GET", testedAddr, nil, func(int, []byte) error {
					address = testedAddr
					return nil
				})
			} else {
				address = dnsAddr
			}

			if err == nil {
				var uaaAddress string
				if uaaAddress, err = findUAAURL(address); err == nil {
					var user string
					var password string

					if user, password, err = opsMan.GetCredentials(
						"deployed",
						"products",
						manifest.Name,
						"credentials",
						".uaa.admin_credentials"); err == nil {
						if err == nil {
							appsMan := &appsManAPI{
								baseAPI: baseAPI{
									clientID:   "cf",
									uaaAddress: uaaAddress,
									address:    address,
									user:       user,
									password:   password,
								},
							}

							if err = appsMan.bindURLMethod(func(items ...string) string {
								// to build the url, first we need to make sure the prefix is correct
								return appsMan.address + "/v2/" + strings.Join(items, "/")
							}); err == nil {
								api = appsMan
							}
						}
					}
				}
			}
		}
	} else {
		err = Error("Nil API can not determin appsman API")
	}

	return api, err
}
