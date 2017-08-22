package helpers

// OpsManConfig of the system.
type OpsManConfig struct {

	// Address of the opsMan
	Address  string `json:"address"`

	// User to use
	User     string `json:"user"`

	// Password to use
	Password string `json:"password"`
}
