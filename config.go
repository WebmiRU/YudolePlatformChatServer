package main

import "YudoleChatServer/packages/channel"

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
}
