package goodgame

type MessageData struct {
	ChannelId string `json:"channel_id"`
	UserId    int    `json:"user_id"`
	UserName  string `json:"user_name"`
	Color     string `json:"color"`
	Icon      string `json:"icon"`
	Role      string `json:"role"`
	Mobile    int    `json:"mobile"`
	IsStatus  int    `json:"isStatus"`
	MessageId int64  `json:"message_id"`
	Timestamp int    `json:"timestamp"`
	Text      string `json:"text"`
	Hidden    int    `json:"hidden"`
	Reload    bool   `json:"reload"`
}

type Message struct {
	Type string      `json:"type"`
	Data MessageData `json:"data"`
}

type JoinRequestData struct {
	ChannelId string `json:"channel_id"`
	Hidden    int    `json:"hidden"`
	Mobile    bool   `json:"mobile"`
	Reload    bool   `json:"reload"`
}

type JoinRequest struct {
	Type string          `json:"type"`
	Data JoinRequestData `json:"data"`
}

type SmileData struct {
	Id         string `json:"id"`
	Key        string `json:"key"`
	Level      int    `json:"level"`
	Paid       string `json:"paid"`
	Bind       string `json:"bind"`
	InternalId int    `json:"internal_id"`
	ChannelId  int    `json:"channel_id"`
	Channel    any    `json:"channel"`
	Nickname   string `json:"nickname"`
	Donat      int    `json:"donat"`
	Premium    int    `json:"premium"`
	Animated   int    `json:"animated"`
	Images     struct {
		Small string `json:"small"`
		Big   string `json:"big"`
		Gif   string `json:"gif"`
	} `json:"images"`
}

type Smile struct {
	Animated string
	Static   string
}
