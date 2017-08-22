package models

// OpsManStatus defines the status
type OpsManStatus struct {
	Status []struct {
		JobName string   `json:"job-name"`
		IPs     []string `json:"ips"`
	} `json:"status"`
}
