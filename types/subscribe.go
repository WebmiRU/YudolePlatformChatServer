package types

type Subscribe struct {
	Type   string   `json:"type"`
	Events []string `json:"events"`
}
