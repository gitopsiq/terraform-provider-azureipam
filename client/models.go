package client

// Space represents a space resource in Azure IPAM
type Space struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Desc string `json:"description"`
}
