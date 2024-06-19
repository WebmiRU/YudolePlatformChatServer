package main

import (
	"YudoleChatServer/packages/resource"
	"YudoleChatServer/packages/theme"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func setCorsJsonHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Allow", "GET,POST,PUT,PATCH,DELETE,HEAD,OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,HEAD,OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-type")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	f, err := os.Open(currentDir + "/http/index.html")
	defer f.Close()

	if err != nil {
		w.WriteHeader(500)
		log.Println("Error opening file", err)
		return
	}

	w.WriteHeader(200)
	if _, err := io.Copy(w, f); err != nil {
		log.Println("Error copying file", err)
		return
	}
}

func indexResourceHandler(w http.ResponseWriter, r *http.Request) {
	path := currentDir + "/http/" + r.URL.Path
	f, err := os.Open(path)
	defer f.Close()

	if err != nil {
		w.WriteHeader(500)
		log.Println("Error opening file:", path, err)
		return
	}

	switch strings.ToLower(filepath.Ext(path)) {
	case ".css", ".scss", ".sass":
		w.Header().Set("Content-Type", "text/css")
	case ".js", "javascript":
		w.Header().Set("Content-Type", "text/javascript")
	case ".html", ".htm":
		w.Header().Set("Content-Type", "text/html")
	}

	w.WriteHeader(200)
	if _, err := io.Copy(w, f); err != nil {
		log.Println("Error copying file", err)
		return
	}
}

func modulesIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	model := resource.ModuleIndex{modules}
	resp, _ := json.Marshal(model)

	w.Write(resp)
}

func modulesIdHandler(w http.ResponseWriter, r *http.Request) {
	setCorsJsonHeaders(&w)

	id := mux.Vars(r)["id"]

	if _, ok := modules[id]; !ok {
		w.WriteHeader(404)
		return
	}

	if r.Method == "PUT" {
		var mod resource.Module
		if err := json.NewDecoder(r.Body).Decode(&mod); err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}

		// @TODO Parse and update only VALUE field
		modules[id].Tabs = mod.Payload.Tabs
		if err := modules[id].Save(); err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}

		go func() {
			_, err := modules[id].RestartWait()
			if err != nil {
				fmt.Printf("Module %s restarting error: %s\n", id, err)
			}

			fmt.Printf("Module %s restarted\n", id)
		}()
	}

	model := resource.Module{
		Payload: modules[id],
	}

	resp, _ := json.Marshal(model)

	w.Write(resp)
}

func modulesIdSetAutostartHandler(w http.ResponseWriter, r *http.Request) {
	setCorsJsonHeaders(&w)

	id := mux.Vars(r)["id"]
	state := mux.Vars(r)["state"]

	if _, ok := modules[id]; !ok {
		w.WriteHeader(404)
		return
	}

	switch r.Method {
	case "PUT":
		model := resource.ModuleIndex{Payload: modules}
		resp, _ := json.Marshal(model)

		if state == "1" && modules[id].Autostart {
			fmt.Println("11")
		} else if state == "1" && !modules[id].Autostart {
			fmt.Println("10")
		} else if state == "0" && modules[id].Autostart {
			fmt.Println("01")
		} else if state == "0" && !modules[id].Autostart {
			fmt.Println("00")
		}

		w.Write(resp)

	case "OPTIONS":
		w.WriteHeader(200)

	default:
		w.WriteHeader(404)
	}
}

func modulesIdStartHandler(w http.ResponseWriter, r *http.Request) {
	setCorsJsonHeaders(&w)

	id := mux.Vars(r)["id"]

	if _, ok := modules[id]; !ok {
		w.WriteHeader(404)
		return
	}

	switch r.Method {
	case "POST":
		go func() {
			if err := modules[id].Start(); err != nil {
				log.Println("Module start error", err)
			}

			broadcast(Message{Type: "api/modules/update/index", Payload: modules})
		}()

		//modules[id].Start()

		model := resource.ModuleIndex{Payload: modules}
		resp, _ := json.Marshal(model)

		w.Write(resp)

	case "OPTIONS":
		w.WriteHeader(200)

	default:
		w.WriteHeader(404)
	}
}

func modulesIdStopHandler(w http.ResponseWriter, r *http.Request) {
	setCorsJsonHeaders(&w)

	id := mux.Vars(r)["id"]

	if _, ok := modules[id]; !ok {
		w.WriteHeader(404)
		return
	}

	switch r.Method {
	case "POST":
		go func() {
			if _, err := modules[id].StopWait(); err != nil {
				log.Println("Module stop error", err)
			}

			broadcast(Message{Type: "api/modules/update/index", Payload: modules})
		}()
		//modules[id].Stop()

		model := resource.ModuleIndex{Payload: modules}
		resp, _ := json.Marshal(model)

		w.Write(resp)

	case "OPTIONS":
		w.WriteHeader(200)

	default:
		w.WriteHeader(404)
	}
}

func themesIndexHandler(w http.ResponseWriter, r *http.Request) {
	setCorsJsonHeaders(&w)

	model := resource.ThemeIndex{themes}
	resp, _ := json.Marshal(model)

	w.Write(resp)
}

func themesPageHandler(w http.ResponseWriter, r *http.Request) {
	setCorsJsonHeaders(&w)

	id := mux.Vars(r)["id"]

	if _, ok := themes[id]; !ok {
		w.WriteHeader(404)
		return
	}

	if r.Method == "PUT" {
		var th resource.Theme
		if err := json.NewDecoder(r.Body).Decode(&th); err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}

		// @TODO Parse and update only VALUE field
		themes[id].Tabs = th.Payload.Tabs
		if err := themes[id].Save(); err != nil {
			log.Println(err)
			w.WriteHeader(400)
			return
		}
	}

	model := resource.Theme{
		Payload: themes[id],
	}

	resp, _ := json.Marshal(model)

	w.Write(resp)
}

func channelsIndexHandler(w http.ResponseWriter, r *http.Request) {
	setCorsJsonHeaders(&w)

	model := resource.ChannelIndex{config.Channels}
	resp, _ := json.Marshal(model)

	w.Write(resp)
}

func channelsPageHandler(w http.ResponseWriter, r *http.Request) {
	setCorsJsonHeaders(&w)

	id := mux.Vars(r)["id"]

	if _, ok := config.Channels[id]; !ok {
		w.WriteHeader(404)
		return
	}

	if r.Method == "PUT" {
		var ch resource.Channel
		if err := json.NewDecoder(r.Body).Decode(&ch); err != nil {
			// @TODO break here
		}

		config.Channels[id] = ch.Payload

		cfg, _ := json.MarshalIndent(config, "", "  ")
		os.WriteFile("./config.json", cfg, 0666)

		for _, client := range sseClients {
			if err := client.SendConfig(); err != nil {
				log.Println(err)
			}
		}
	}

	// Remap default theme values with channel settings
	if _, ok := themes[config.Channels[id].Theme.Name]; ok {
		if _, ok := config.Channels[id].Theme.Config[config.Channels[id].Theme.Name]; ok {
			for tabKey, tabVal := range themes[config.Channels[id].Theme.Name].Tabs {
				for field, attrs := range tabVal.Fields {
					if _, ok := config.Channels[id].Theme.Config[config.Channels[id].Theme.Name].Tabs[tabKey].Fields[field]; ok {
						newAttrs := attrs
						newAttrs.Value = config.Channels[id].Theme.Config[config.Channels[id].Theme.Name].Tabs[tabKey].Fields[field].Value
						themes[config.Channels[id].Theme.Name].Tabs[tabKey].Fields[field] = newAttrs
					}
				}
			}
		}
	}

	payload := config.Channels[id]
	// @TODO Сделать проверку
	payload.Theme.Config = make(map[string]*theme.Theme)
	payload.Theme.Config[config.Channels[id].Theme.Name] = themes[config.Channels[id].Theme.Name]

	if config.Channels[id].ServiceIcons == nil {
		config.Channels[id].ServiceIcons = make(map[string]string)
	}

	for _, v := range services {
		if _, ok := config.Channels[id].ServiceIcons[v]; !ok {
			config.Channels[id].ServiceIcons[v] = ""
		}
	}

	fmt.Println(config.Channels[id].ServiceIcons)

	model := resource.Channel{
		Payload: payload,
	}

	resp, _ := json.Marshal(model)

	w.Write(resp)
}
