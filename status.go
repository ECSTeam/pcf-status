package main

// Status of this platform.
type Status struct {
	Error         string `json:"error,omitempty"`
	Legacy        bool   `json:"legacy"`
	OpsManVersion string `json:"opsman-version,omitempty"`
	ErtVersion    string `json:"ert-version,omitempty"`
}

var opsManClient *OpsManClient

const (
	legacyVersionString = "\u2264 1.6"

	// UnknownVersionString holds the unknown version string.
	UnknownVersionString = "unknown"
)

// NewStatus will create a new status object.
func NewStatus() (status Status, err error) {

	if opsManClient == nil {
		if opsManClient, err = NewOpsManClient(); err != nil {
			return
		}
	}

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

		status = Status{
			OpsManVersion: opsManVer,
			ErtVersion:    ertVer,
			Legacy:        info.Legacy,
		}
	}

	return
}
