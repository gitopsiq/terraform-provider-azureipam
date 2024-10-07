package client

// Space represents an IPAM space
type Space struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Block struct {
    ID   string `json:"id"`
    CIDR string `json:"cidr"`
}

type Reservation struct {
    ID   string `json:"id"`
    CIDR string `json:"cidr"`
}