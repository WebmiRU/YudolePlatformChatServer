package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"slices"
)

func main() {
	go tcpServerStart()
	wsServerStart()
}

func broadcast(message json.RawMessage, event string) {
	log.Println("BROADCAST")
	for socket, events := range tcpSubscribers {
		if slices.Index(events, event) >= 0 {
			socket.Write(message)
		}
	}

	for socket, events := range wsSubscribers {
		if slices.Index(events, event) >= 0 {
			socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
