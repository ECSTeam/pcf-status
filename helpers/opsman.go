package helpers

import "strings"

// opsManAPI is the structure for the OpsMan interface.
type opsManAPI struct {
	baseAPI
}

// NewOpsManAPI creates a new api.
func NewOpsManAPI(uaaAddress string, address string, user string, password string) (API, error) {
	opsMan := &opsManAPI{
		baseAPI: baseAPI{
			clientID:   "opsman",
			uaaAddress: uaaAddress,
			address:    address,
			user:       user,
			password:   password,
		},
	}

	err := opsMan.bindURLMethod(func(items ...string) string {
		// to build the url, first we need to make sure the prefix is correct
		return address + "/api/v0/" + strings.Join(items, "/")
	})

	return opsMan, err
}
