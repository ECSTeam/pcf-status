package helpers

import "strings"

// appsManAPI is the structure for the AppsMan interface.
type appsManAPI struct {
	baseAPI
}

// NewAppsManAPI creates a new api.
func NewAppsManAPI(uaaAddress string, address string, user string, password string) (API, error) {
	appsMan := &appsManAPI{
		baseAPI: baseAPI{
			clientID:   "cf",
			uaaAddress: uaaAddress,
			address:    address,
			user:       user,
			password:   password,
		},
	}

	err := appsMan.bindURLMethod(func(items ...string) string {
		// to build the url, first we need to make sure the prefix is correct
		return address + "/v2/" + strings.Join(items, "/")
	})

	return appsMan, err
}
