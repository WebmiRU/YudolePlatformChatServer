package channel

import (
	"YudoleChatServer/packages/theme"
)

type Channel struct {
	Title        string            `json:"title"`
	Events       []string          `json:"events"`
	ServiceIcons map[string]string `json:"service_icons"`

	Theme struct {
		Name string `json:"name"`
		//Values map[string]map[string]any `json:"values"`
		Config map[string]*theme.Theme `json:"config"`
	} `json:"theme"`
}
