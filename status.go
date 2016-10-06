package main

import "errors"

// Version of the item.
type Version struct {
	Version  string `json:"ver,omitempty"`
	StemCell string `json:"sc,omitempty"`
}

// Status of this platform.
type Status struct {
	Error    string             `json:"error,omitempty"`
	Versions map[string]Version `json:"versions,omitempty"`
}

const (
	legacyVersionString = "\u2264 1.6"

	// UnknownVersionString holds the unknown version string.
	UnknownVersionString = "unknown"

	// These are the labels for opsman and ERT. They have the special
	// characters to ensure they are sorted first.
	nameOpsMan = "\001Ops Man"
	nameErt    = "\002ERT"
)

// StandardNames are the mappings between labels and nice names.
// NOTE: Add tile translations here: if the value is empty, then that
//       entry will be ignored.
var StandardNames = map[string]string{
	"cf":      nameErt,
	"p-bosh":  "",
	"p-mysql": "MySql Tile",
}

// NewStatus will create a new status object.
func NewStatus(includes *Includes) (status *Status, err error) {

	var opsManClient *OpsManClient
	if opsManClient, err = NewOpsManClient(); err == nil {
		info := DiagnosticReport{}
		if err = opsManClient.GetInfo(&info); err == nil {

			if !info.Legacy {
				status = &Status{
					Versions: map[string]Version{
						nameOpsMan: Version{
							Version: info.Versions.Release,
						},
					},
				}

				for _, item := range info.Products.Deployed {

					stemcell := ""
					if includes.StemCellVersion {
						stemcell = item.Stemcell
					}

					if name, ok := StandardNames[item.Name]; ok {

						// Ignore empty values.
						if len(name) > 0 {
							status.Versions[name] = Version{
								Version:  item.Version,
								StemCell: stemcell,
							}
						}
					} else {
						status.Versions[item.Name] = Version{
							Version:  item.Version,
							StemCell: stemcell,
						}
					}
				}
			} else {
				status = &Status{
					Versions: map[string]Version{
						nameOpsMan: {Version: legacyVersionString},
						nameErt:    {Version: legacyVersionString},
					},
				}
			}
		}
	}

	if err != nil {
		status = &Status{
			Error: err.Error(),
		}
	} else {
		if status == nil {
			err = errors.New("Unknown status")
		} else if len(status.Versions) == 0 {
			err = errors.New("No versions found")
		}
	}

	return
}
