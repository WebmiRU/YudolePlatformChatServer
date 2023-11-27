package types

type Unsubscribe struct {
	Type   string   `json:"type"`
	Events []string `json:"events"`
}
