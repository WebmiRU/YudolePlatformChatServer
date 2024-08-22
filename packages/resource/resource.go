package resource

type Audio struct {
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	Sha256   string `json:"sha256"`
	MimeType string `json:"mime_type"`
	Source   string `json:"source"`
	Path     string `json:"-"`
}

type Image struct {
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	Sha256   string `json:"sha256"`
	MimeType string `json:"mime_type"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Source   string `json:"source"`
	Path     string `json:"-"`
}

type Payload struct {
	Audio []Audio `json:"audio"`
	Image []Image `json:"image"`
}

type Index struct {
	Type    string  `json:"type"` // "resources/audio"
	Payload Payload `json:"payload"`
}
