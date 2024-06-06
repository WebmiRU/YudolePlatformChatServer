package tab

type Item struct {
	Title string `json:"title"`
	Value any    `json:"value"`
}

type Field struct {
	Type        string `json:"type"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Placeholder string `json:"placeholder"`
	Validation  string `json:"validation"`
	Value       any    `json:"value"`
	Items       []Item `json:"items"`
}

type Tab struct {
	Title  string           `json:"title"`
	Fields map[string]Field `json:"fields"`
}
