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

func wsServerStart() {
	http.HandleFunc("/chat", wsAccept)
	log.Fatal(http.ListenAndServe("0.0.0.0:5800", nil))
}
