package main

import "errors"

// Status of this platform.
type Status struct {
	Error    string            `json:"error,omitempty"`
	Versions map[string]string `json:"versions,omitempty"`
}

const (
	legacyVersionString = "\u2264 1.6"

	// UnknownVersionString holds the unknown version string.
	UnknownVersionString = "unknown"
)

// NewStatus will create a new status object.
func NewStatus() (status *Status, err error) {

	var opsManClient *OpsManClient
	if opsManClient, err = NewOpsManClient(); err == nil {
		info := DiagnosticReport{}
		if err = opsManClient.GetInfo(&info); err == nil {

			opsManVer := UnknownVersionString
			ertVer := UnknownVersionString
			if !info.Legacy {
				opsManVer = info.Versions.Meta
				for _, item := range info.Products.Deployed {
					if item.Name == "cf" {
						ertVer = item.Version
						break
					}
				}
			} else {
				opsManVer = legacyVersionString
				ertVer = legacyVersionString
			}

			status = &Status{
				Versions: map[string]string{
					"Ops Man": opsManVer,
					"ERT":     ertVer,
				},
			}
		}
	}

	if status == nil {
		err = errors.New("Unknown status")
	} else if len(status.Versions) == 0 {
		err = errors.New("No versions found")
	}

	if err != nil {
		status = &Status{
			Error: err.Error(),
		}
	}

	return
}
