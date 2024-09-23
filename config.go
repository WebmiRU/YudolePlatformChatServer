package main

import (
	"YudoleChatServer/packages/channel"
	"YudoleChatServer/packages/resource"
	"encoding/json"
	"log"
	"os"
)

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
	config.Resources.Audio = make(map[string]resource.Audio)
	//config.Resources.Images = []resource.Image

	configFile, err := os.Open("config.json")

	if err != nil {
		panic("Error while reading \"config.json\" file")
	}

	if err := json.NewDecoder(configFile).Decode(config); err != nil {
		panic("Error while parsing \"config.json\" file")
	}

	//resources.Load()
}

func (config *Config) Save() {
	file, _ := json.MarshalIndent(config, "", "  ")
	err := os.WriteFile("config.json", file, 0644)

	if err != nil {
		log.Println("Error while writing \"config.json\" file")
	}

	//resources.Load()
}
