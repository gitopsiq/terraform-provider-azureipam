package client

// Space represents a space in the Azure IPAM system
type Space struct {
	ID   string `json:"id,omitempty"`   // ID is the unique identifier for the space
	Name string `json:"name"`           // Name is the name of the space
	Desc string `json:"desc,omitempty"` // Desc is an optional description of the space
}
