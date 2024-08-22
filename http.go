package main

import (
	"YudoleChatServer/packages/resource"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func response(w http.ResponseWriter, t string, data any) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	resp1 := resource.Response{
		Type:    t,
		Payload: data,
	}

	resp, _ := json.Marshal(resp1)
	defer w.Write(resp)
}

func FileGetResourceAudio(path string) (*resource.Audio, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	s256 := sha256.New()
	io.Copy(s256, file)
	sha256String := hex.EncodeToString(s256.Sum(nil))

	stat, _ := file.Stat()
	buff := make([]byte, 100)
	file.ReadAt(buff, 0)
	mimeType := http.DetectContentType(buff)

	result := resource.Audio{
		Name:     stat.Name(),
		Size:     stat.Size(),
		Sha256:   sha256String,
		MimeType: mimeType,
	}

	file.Close()

	return &result, nil
}

func FileGetResourceImage(path string) (*resource.Image, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	s256 := sha256.New()
	io.Copy(s256, file)
	sha256String := hex.EncodeToString(s256.Sum(nil))

	stat, _ := file.Stat()
	buff := make([]byte, 100)
	file.ReadAt(buff, 0)

	name := stat.Name()
	size := stat.Size()
	mimeType := http.DetectContentType(buff)

	file.Close()

	imageFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	im, _, err := image.DecodeConfig(imageFile)

	width := im.Width
	height := im.Height

	imageFile.Close()

	result := resource.Image{
		Name:     name,
		Size:     size,
		Sha256:   sha256String,
		MimeType: mimeType,
		Width:    width,
		Height:   height,
	}

	return &result, nil
}

func resourcesAudioPayload() map[string]resource.Audio {
	payload := make(map[string]resource.Audio)

	for _, v := range config.Resources.Audio {
		payload[v.Sha256] = v
	}

	for _, module := range modules {
		for _, v := range module.Resources.Audio {
			res, err := FileGetResourceAudio(module.Dir + "/resources/" + v)

			if err != nil {
				log.Println(err)
				continue
			}

			payload[res.Sha256] = *res
		}
	}

	return payload
}

func resourcesImagesPayload() map[string]resource.Image {
	payload := make(map[string]resource.Image)

	for _, v := range config.Resources.Images {
		payload[v.Sha256] = v
	}

	for _, module := range modules {
		for _, v := range module.Resources.Images {
			res, err := FileGetResourceImage(module.Dir + "/resources/" + v)

			if err != nil {
				log.Println(err)
				continue
			}

			payload[res.Sha256] = *res
		}
	}

	return payload
}

func resourcesAudioGet(w http.ResponseWriter, r *http.Request) {
	response(w, "resources/audio", resourcesAudioPayload())
}

func resourcesAudioPost(w http.ResponseWriter, r *http.Request) {
	defer response(w, "resources/audio", &config.Resources.Audio)

	httpFile, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error retrieving the File", err)
		return
	}

	tmpFile, err := os.CreateTemp("", "tmp_")

	if err != nil {
		log.Println("Error creating a temporary file", err)
		return
	}

	io.Copy(tmpFile, httpFile)

	tmpFile.Close()

	file, err := os.Open(tmpFile.Name())
	if err != nil {
		log.Println("Error opening the file", err)
	}

	s256 := sha256.New()
	io.Copy(s256, file)
	sha256String := hex.EncodeToString(s256.Sum(nil))
	file.Close()

	dstFilePath := "data/resources/audio/" + sha256String
	dstFile, err := os.Create(dstFilePath)
	if err != nil {
		log.Println("Error creating the file", err)
		return
	}

	file2, _ := os.Open(tmpFile.Name())

	io.Copy(dstFile, file2)
	dstFile.Close()

	buff := make([]byte, 100)
	file2.ReadAt(buff, 0)
	mimeType := http.DetectContentType(buff)

	file2.Close()

	config.Resources.Audio[sha256String] = resource.Audio{
		Name:     handler.Filename,
		Sha256:   sha256String,
		Size:     handler.Size,
		MimeType: mimeType,
	}

	config.Save()
}

func resourcesImagesGet(w http.ResponseWriter, r *http.Request) {
	response(w, "resources/images", resourcesImagesPayload())
}

func resourcesIndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// @TODO
	//model := resource.ResourceIndex{resources}
	//resp, _ := json.Marshal(model)
	//
	//w.Write(resp)
}

func httpResource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var path string

	switch vars["type"] {
	case "module":
		allow := false

		if _, ok := modules[vars["id"]]; ok {
			//for _, res := range modules[vars["id"]].Resources {
			//	if res.Path == vars["path"] {
			//		allow = true
			//	}
			//}
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

	// Audio resources
	router.HandleFunc("/api/resources/audio", resourcesAudioGet).Methods("GET")
	router.HandleFunc("/api/resources/audio", resourcesAudioPost).Methods("POST")

	// Image resources
	router.HandleFunc("/api/resources/images", resourcesImagesGet).Methods("GET")

	http.Handle("/", router)

	fmt.Println("Server starting...")
	http.ListenAndServe(":80", nil) // 5128
}
