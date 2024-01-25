package obj

type GoodGame struct {
	Login    string  `json:"login"`
	Password string  `json:"password"`
	Channels []int64 `json:"channels"`
}

type Trovo struct {
	Login    string  `json:"login"`
	Password string  `json:"password"`
	Channels []int64 `json:"channels"`
}

type Twitch struct {
	Login    string  `json:"login"`
	Password string  `json:"password"`
	Channels []int64 `json:"channels"`
}

type YouTube struct {
	Login    string  `json:"login"`
	Password string  `json:"password"`
	Channels []int64 `json:"channels"`
}

type Services struct {
	GoodGame GoodGame `json:"goodgame"`
	Trovo    Trovo    `json:"trovo"`
	Twitch   Twitch   `json:"twitch"`
	YouTube  YouTube  `json:"youtube"`
}

type Config struct {
	Services Services `json:"services"`
}
