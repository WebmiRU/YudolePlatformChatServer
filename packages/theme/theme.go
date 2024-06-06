package theme

import (
	"encoding/json"
	"os"
)

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

type Theme struct {
	Title string         `json:"title"`
	Tabs  map[string]Tab `json:"tabs"`

	dir        string `json:"dir"`
	configPath string
}

func (m *Theme) Load(configPath string) error {
	m.configPath = configPath + string(os.PathSeparator) + "theme.json"
	configBytes, _ := os.ReadFile(m.configPath)
	m.dir = configPath

	if err := json.Unmarshal(configBytes, &m); err != nil {
		return err
	}

	return nil
}

func (m *Theme) Save() error {
	if data, err := json.MarshalIndent(m, "", "    "); err != nil {
		return err
	} else {
		if err := os.WriteFile(m.configPath, data, 0666); err != nil {
			return err
		}
	}

	return nil
}
