package types

type UserMeta struct {
	Avatar string `json:"avatar"`
	Badges string `json:"badges"`
}

type User struct {
	Nickname string   `json:"nickname"`
	Login    string   `json:"login"`
	Meta     UserMeta `json:"meta"`
}

type ChatMessage struct {
	Type string `json:"type"`
	Src  string `json:"src"`
	Text string `json:"text"`
	Html string `json:"html"`
	User User   `json:"user"`
}
