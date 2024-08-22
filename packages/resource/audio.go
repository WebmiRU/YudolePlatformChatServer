package resource

type Audio struct {
	Type    string                   `json:"type"`
	Payload map[string]AudioResource `json:"payload"`
}

type AudioPayload struct {
	Payload map[string]AudioResource `json:"payload"`
}

type AudioResource struct {
	Name     string `json:"name"`
	Sha256   string `json:"sha256"`
	Size     int    `json:"size"`
	MimeType string `json:"mime_type"`
}
