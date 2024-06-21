package main

import (
	"YudoleChatServer/packages/resource"
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

func resourcesIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	model := resource.ResourceIndex{resources}
	resp, _ := json.Marshal(model)

	w.Write(resp)
}

func httpResource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var path string

	switch vars["type"] {
	case "module":
		allow := false

		if _, ok := modules[vars["id"]]; ok {
			for _, res := range modules[vars["id"]].Resources {
				if res.Path == vars["path"] {
					allow = true
				}
			}
		}

		if !allow {
			w.WriteHeader(404)
			return
		}

		path = currentDir + "/modules/" + vars["id"] + "/resources/" + vars["path"]
	case "theme":
		path = currentDir + "/themes/" + vars["id"] + "/resources/" + vars["path"]
	}

	fmt.Println(path)

	f, err := os.Open(path)
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
	//w.Header().Set("Content-Type", "image/*")
	//w.Header().Set("Content-Length", len)
}

//func httpFile(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//
//	var path string
//
//	switch vars["type"] {
//	case "module":
//		allow := false
//
//	loop:
//		for _, resType := range resources {
//			for _, resource := range resType {
//				if resource == vars["path"] {
//					allow = true
//					break loop
//				}
//			}
//		}
//
//		if !allow {
//			w.WriteHeader(404)
//		}
//
//		path = currentDir + "/modules/" + vars["id"] + "/resources/" + vars["path"]
//	case "theme":
//		path = currentDir + "/themes/" + vars["id"] + "/resources/" + vars["path"]
//	}
//
//	fmt.Println(path)
//
//	f, err := os.Open(path)
//	defer f.Close()
//
//	if err != nil {
//		w.WriteHeader(500)
//		log.Println("Error opening file", err)
//		return
//	}
//
//	w.WriteHeader(200)
//	if _, err := io.Copy(w, f); err != nil {
//		log.Println("Error copying file", err)
//		return
//	}
//	//w.Header().Set("Content-Type", "image/*")
//	//w.Header().Set("Content-Length", len)
//}

func httpChat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	w.Write([]byte(vars["channel"]))
}

func httpChatChannelHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	channel := vars["channel"]
	path := vars["path"]

	if _, ok := config.Channels[channel]; !ok {
		w.Write([]byte("Channel Not Found"))
		return
	}

	theme := config.Channels[channel].Theme.Name

	//var f *os.File
	//var err error

	var fullPath string

	if len(path) > 0 {
		fullPath = "./themes/" + theme + "/" + path
	} else {
		fullPath = "./themes/" + theme + "/index.html"
	}

	f, err := os.Open(fullPath)

	switch strings.ToLower(filepath.Ext(path)) {
	case ".css": //, ".scss", ".sass":
		w.Header().Set("Content-Type", "text/css")
	case ".scss", ".sass":
		w.Header().Set("Content-Type", "text/css")
	case ".js", "javascript":
		w.Header().Set("Content-Type", "text/javascript")
	case ".html", ".htm":
		w.Header().Set("Content-Type", "text/html")
	}

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

func httpServer() {
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/assets/{path.*}", indexResourceHandler)
	router.HandleFunc("/chat/{channel}/", httpChatChannelHandler).Methods("GET")
	router.HandleFunc("/chat/{channel}/{path:.*}", httpChatChannelHandler).Methods("GET")
	router.HandleFunc("/chat2/{channel}", httpChat)
	router.HandleFunc("/sse/{channel}", sseChannelGet).Methods("GET")

	// API
	router.HandleFunc("/api/events", apiEventsHandlerGet).Methods("GET")
	router.HandleFunc("/api/themes", themesIndexHandler)
	router.HandleFunc("/api/themes/{id}", themesPageHandler)
	router.HandleFunc("/api/channels", channelsIndexHandler)
	router.HandleFunc("/api/channels/{id}", channelsPageHandler)
	router.HandleFunc("/api/modules", modulesIndexHandler)
	router.HandleFunc("/api/modules/{id}", modulesIdHandler)
	router.HandleFunc("/api/modules/{id}/start", modulesIdStartHandler)
	router.HandleFunc("/api/modules/{id}/stop", modulesIdStopHandler)
	router.HandleFunc("/api/modules/{id}/autostart/{state:[0,1]}", modulesIdSetAutostartHandler)
	router.HandleFunc("/api/resources", resourcesIndexHandler)
	router.HandleFunc("/resource/{type:module|theme}/{id}/{path:.*}", httpResource).Methods("GET")
	http.Handle("/", router)

	fmt.Println("Server starting...")
	http.ListenAndServe(":80", nil) // 5128
}
