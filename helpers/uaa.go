package helpers

import (
	"encoding/json"
	"log"
	"net/http"
)

// opsManBaseResponse get the base response
type baseResponse struct {
	Links map[string]interface{} `json:"links"`
}

// findUAAURL finds the uaa address from the root.
func findUAAURL(address string) (uaa string, err error) {
	err = callURL("GET", address, nil, func(status int, data []byte) (e error) {

		if status >= http.StatusOK && status < http.StatusBadRequest {

			var resp baseResponse
			if e = json.Unmarshal(data, &resp); e == nil {
				var ok bool
				var uaaLinks interface{}
				if uaaLinks, ok = resp.Links["uaa"]; ok {

					var uaaStruct map[string]interface{}
					if uaaStruct, ok = uaaLinks.(map[string]interface{}); ok {
						var has bool
						if uaaLinks, has = uaaStruct["href"]; !has {
							e = Error("No href in uaa structure.")
						}
					}

					if e == nil {
						if uaa, ok = uaaLinks.(string); ok {
							log.Printf("Found UAA: %s", uaa)
						}
					}
				} else {
					e = Error("Unable to find UAA Address")
				}
			}
		} else {
			e = Error("Unable to access OpsMan endpoint.")
		}

		return e
	})

	return uaa, err
}
