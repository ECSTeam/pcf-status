package main

// Status of this platform.
type Status struct {

	// Info is the
	OpsManVersion string `json:"opsman-version"`
	ErtVersion    string `json:"ert-version"`
}

var opsManClient *OpsManClient

// NewStatus will create a new status object.
func NewStatus() (status Status, err error) {

	if opsManClient == nil {
		if opsManClient, err = NewOpsManClient(); err != nil {
			return
		}
	}

	info := DiagnosticReport{}
	if err = opsManClient.GetInfo(&info); err == nil {

		ertVer := "unknown"
		for _, item := range info.Products.Deployed {
			if item.Name == "cf" {
				ertVer = item.Version
				break
			}
		}

		status = Status{
			OpsManVersion: info.Versions.Meta,
			ErtVersion:    ertVer,
		}
	}

	return
}
