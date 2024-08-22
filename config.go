package main

import (
	"YudoleChatServer/packages/channel"
	"encoding/json"
	"log"
	"os"
)

type Resource struct {
	Name     string `json:"name"`
	Sha256   string `json:"sha256"`
	Size     int64  `json:"size"`
	MimeType string `json:"mime_type"`
}

type Resources struct {
	Audio  map[string]Resource
	Images map[string]Resource
}

type Config struct {
	AutostartModules []string `json:"autostart_modules"`

	TcpServer struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"tcp_server"`

	HttpServer struct {
		Host string `json:"host"`
		Port int    `json:"port"`
	} `json:"http_server"`

	Channels map[string]*channel.Channel `json:"channels"`

	Resources Resources `json:"resources"`
}

func (config *Config) Load() {
	config.Resources.Audio = make(map[string]Resource)
	config.Resources.Images = make(map[string]Resource)

	configFile, err := os.Open("config.json")

	if err != nil {
		panic("Error while reading \"config.json\" file")
	}

	if err := json.NewDecoder(configFile).Decode(config); err != nil {
		panic("Error while parsing \"config.json\" file")
	}
}

func (config *Config) Save() {
	file, _ := json.MarshalIndent(config, "", "  ")
	err := os.WriteFile("config.json", file, 0644)

	if err != nil {
		log.Println("Error while writing \"config.json\" file")
	}
}
