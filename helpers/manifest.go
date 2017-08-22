package helpers

// Stemcell defines the types of stemcells available.
type Stemcell struct {
	Source  string `json:"source,omitempty"`
	Alias   string `json:"alias"`
	OS      string `json:"os"`
	Version string `json:"version"`
}

// Manifest defines the manifest
type Manifest struct {
	Name          string     `json:"name"`
	Stemcells     []Stemcell `json:"stemcells"`
	InstanceGroup []struct {
		Name       string                 `json:"name"`
		Properties map[string]interface{} `json:"properties"`

		AvaliabilityZone []string `json:"azs"`
		Instances        int      `json:"instances"`
		Lifecycle        string   `json:"lifecycle"`
		VMType           string   `json:"vm_type"`
		Stemcell         string   `json:"stemcell"`
		Networks         []struct {
			Name      string   `json:"name"`
			StaticIPs []string `json:"static_ips"`
		} `json:"networks"`
	} `json:"instance_groups"`
}
