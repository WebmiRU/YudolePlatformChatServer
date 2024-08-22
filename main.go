package main

import (
	"YudoleChatServer/packages/module"
	"YudoleChatServer/packages/theme"
	"fmt"
	"log"
	"os"
	"os/signal"
	"slices"
	"strings"
	"syscall"
)

type IMessage interface {
	GetType() string
}

type Message struct {
	Id      string `json:"id"`
	Module  string `json:"module"`
	Type    string `json:"type"`
	Payload any    `json:"payload"`
}

func (m Message) GetType() string {
	return m.Type
}

var signals = make(chan os.Signal, 99)
var config Config
var currentDir string
var modules = make(map[string]*module.Module)
var themes = make(map[string]*theme.Theme)
var services []string
var resources map[string][]string
var events = []string{
	//"event/subscribe",
	//"event/unsubscribe",
	"stream/chat/message",
	"stream/chat/private_message",
	"api/modules/update/index",
} // All known events

func broadcast(message IMessage) {
	//fmt.Println("BROADCAST", message.GetType(), sseClients)

	// Broadcast SSE clients
	for _, client := range sseClients {
		fmt.Println(client.Events, message.GetType())
		if slices.Contains(client.Events, message.GetType()) {
			err := client.Send(message)
			if err != nil {
				log.Println("Error sending message:", err)
			}
		}
	}

	// Broadcast TCP clients
	for _, client := range tcpClients {
		if slices.Contains(client.Events, message.GetType()) {
			err := client.Send(message)
			if err != nil {
				log.Println("Error sending message:", err)
			}
		}
	}
}

func Init() {
	// Catch shutdown signals from OS
	signal.Notify(signals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)
	go shutdown()

	//// Loading local config file
	//configFile, err := os.Open("config.json")
	//
	//if err != nil {
	//	panic("Error while reading \"config.json\" file")
	//}
	//
	//if err := json.NewDecoder(configFile).Decode(&config); err != nil {
	//	panic("Error while parsing \"config.json\" file")
	//}

	config.Load()

	// Run TCP server
	go tcpServer()
}

func shutdown() {
	for {
		select {
		case <-signals:
			for _, m := range modules {
				m.Stop()
			}

			log.Println("Shutting down...")
			os.Exit(0)
		}
	}
}

func loadModules() {
	resources = make(map[string][]string)
	moduleList, _ := os.ReadDir(currentDir + fmt.Sprintf("%c%s", os.PathSeparator, "modules"))

	for _, dir := range moduleList {
		path := currentDir + string(os.PathSeparator) + "modules" + string(os.PathSeparator) + dir.Name()

		var mod module.Module
		if err := mod.Load(path); err == nil {
			// Если модуль имеет тип "Клиент" - добавляем его в список сервисов
			_type := strings.ToLower(mod.Type)
			_service := strings.ToLower(mod.Service)
			if _type == "client" && !slices.Contains(services, _service) {
				services = append(services, _service)
			}

			for _, v := range mod.Resources {
				resources[v.Type] = append(resources[v.Type], "/resource/module/"+dir.Name()+"/"+v.Path)
			}

			// Добавляем модуль в список модулей
			modules[dir.Name()] = &mod
		} else {
			log.Println(err)
			continue
		}

		if err := mod.Start(); err != nil {
			log.Println(err)
			continue
		}
	}

	//fmt.Println(resources)

	loadThemes()
}

func loadThemes() {
	themesList, _ := os.ReadDir(currentDir + fmt.Sprintf("%c%s", os.PathSeparator, "themes"))
	//log.Println("THL", themesList)

	for _, dir := range themesList {
		path := currentDir + string(os.PathSeparator) + "themes" + string(os.PathSeparator) + dir.Name()

		var th theme.Theme
		if err := th.Load(path); err == nil {
			themes[dir.Name()] = &th
		} else {
			log.Println(err)
		}
	}
}

func main() {
	currentDir, _ = os.Getwd()

	Init()
	loadModules()

	//go func() {
	//	for {
	//		time.Sleep(5 * time.Second)
	//		for _, c := range sseClients {
	//			c.Send(Message{Module: "client", Type: "api/modules/update/index", Payload: nil})
	//		}
	//		//fmt.Println("RESTART")
	//		//code, err := modules["twitch_client"].RestartWait()
	//		//fmt.Println(code, err)
	//		//fmt.Println("STARTED")
	//	}
	//}()

	httpServer()
}
