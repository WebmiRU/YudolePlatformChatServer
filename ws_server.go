package main

import (
	"YudolePlatofrmChatServer/obj"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	//ReadBufferSize:  4096,
	//WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsAccept(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	defer conn.Close()

	if err != nil {
		log.Println("Websocket upgrading error:", err)
		return
	}

	write := func(message json.RawMessage) error {
		return conn.WriteJSON(message)
	}

	client := obj.Client{
		Write: write,
	}

	obj.ClientAppend(&client)

	for {
		var message json.RawMessage
		if err := conn.ReadJSON(&message); err != nil {
			log.Println(err)
			break
		}
	}

	obj.ClientRemove(&client)
}

func configResponse(w http.ResponseWriter, r *http.Request) {
	cfg, _ := json.Marshal(config)

	log.Println(r.Method)
	switch r.Method {
	case "GET":
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(cfg)
		break

	case "POST":
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)

		jd := json.NewDecoder(r.Body)
		jd.DisallowUnknownFields()
		jd.Decode(&config)

		configSave()
		break

	case "OPTIONS":
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(200)

	default:
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.WriteHeader(404)
	}

	return
}

func wsServerStart() {
	http.HandleFunc("/config", configResponse)
	http.HandleFunc("/chat", wsAccept)
	log.Fatal(http.ListenAndServe("0.0.0.0:5800", nil))
}
