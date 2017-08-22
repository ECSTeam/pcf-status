package helpers

// DiagnosticReport returns the diagnostic report
type DiagnosticReport struct {
	Versions struct {
		Release string `json:"release_version"`
	} `json:"versions"`

	Type string `json:"infrastructure_type"`

	Releases []string `json:"releases"`

	Stemcells []string `json:"stemcells"`

	Products struct {
		Deployed []struct {
			Name     string `json:"name"`
			Version  string `json:"version"`
			Stemcell string `json:"stemcell"`
		} `json:"deployed"`
	}
}
