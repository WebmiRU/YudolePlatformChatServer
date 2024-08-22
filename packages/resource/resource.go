package resource

type IResource interface {
	GetPath() string
	GetMimeType() string
	GetSize() int64
}

type Audio struct {
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	Sha256   string `json:"sha256"`
	MimeType string `json:"mime_type"`
	Source   string `json:"source"`
	Path     string `json:"-"`
}

func (a *Audio) GetPath() string {
	return a.Path
}

func (a *Audio) GetMimeType() string {
	return a.MimeType
}

func (a *Audio) GetSize() int64 {
	return a.Size
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

func (i *Image) GetPath() string {
	return i.Path
}

func (i *Image) GetMimeType() string {
	return i.MimeType
}

func (i *Image) GetSize() int64 {
	return i.Size
}

type Payload struct {
	Audio []Audio `json:"audio"`
	Image []Image `json:"image"`
}

type Index struct {
	Type    string  `json:"type"` // "resources/audio"
	Payload Payload `json:"payload"`
}
