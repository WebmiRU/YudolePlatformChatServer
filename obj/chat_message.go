package obj

type Base struct {
	Type string `json:"type"`
}

type UserMeta struct {
	Avatar string            `json:"avatar"`
	Badges map[string]string `json:"badges"`
}

type User struct {
	Name string   `json:"nickname"`
	Meta UserMeta `json:"meta"`
}

type ChatMessage struct {
	Type    string `json:"type"`
	Service string `json:"service"`
	Src     string `json:"src"`
	Text    string `json:"text"`
	Html    string `json:"html"`
	User    User   `json:"user"`
}

type Subscribe struct {
	Type   string   `json:"type"`
	Events []string `json:"events"`
}

type Unsubscribe struct {
	Type   string   `json:"type"`
	Events []string `json:"events"`
}
