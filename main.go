package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"slices"
)

func main() {
	go tcpServerStart()
	wsServerStart()
}

func broadcast(message json.RawMessage, event string) {
	tcpSubscribersMutex.Lock()
	for socket, events := range tcpSubscribers {
		if slices.Index(events, event) >= 0 {
			socket.Write(message)
		}
	}
	tcpSubscribersMutex.Unlock()

	wsSubscribersMutex.Lock()
	for socket, events := range wsSubscribers {
		if slices.Index(events, event) >= 0 {
			socket.WriteMessage(websocket.TextMessage, message)
		}
	}
	wsSubscribersMutex.Unlock()
}
