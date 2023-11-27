package main

import (
	"YudolePlatofrmChatServer/types"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"slices"
	"sync"
)

var wsSubscribersMutex sync.Mutex
var wsSubscribers = make(map[*websocket.Conn][]string)

var upgrader = websocket.Upgrader{
	//ReadBufferSize:  4096,
	//WriteBufferSize: 4096,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func wsAccept(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Websocket upgrading error:", err)
		return
	}

	defer conn.Close()

	for {
		var message json.RawMessage
		if err := conn.ReadJSON(&message); err != nil {
			break
		}

		if err != nil {
			log.Println("JSON message decode error:", err)
			continue
		}

		var base types.Base
		if err := json.Unmarshal(message, &base); err != nil {
			log.Println("JSON message decode error:", err)
			continue
		}

		switch base.Type {
		case "subscribe":
			var subscribe types.Subscribe
			if err := json.Unmarshal(message, &subscribe); err != nil {
				log.Println("JSON message decode error:", err)
				continue
			}

			wsSubscribersMutex.Lock()
			for _, v := range subscribe.Events {
				if slices.Index(wsSubscribers[conn], v) == -1 {
					wsSubscribers[conn] = append(wsSubscribers[conn], v)
				}
			}
			wsSubscribersMutex.Unlock()
			break

		case "unsubscribe":
			var unsubscribe types.Unsubscribe
			if err := json.Unmarshal(message, &unsubscribe); err != nil {
				log.Println("JSON message decode error:", err)
				continue
			}

			wsSubscribersMutex.Lock()
			for _, v := range unsubscribe.Events {
				idx := slices.Index(wsSubscribers[conn], v)

				if slices.Index(wsSubscribers[conn], v) >= 0 {
					wsSubscribers[conn] = slices.Delete(wsSubscribers[conn], idx, idx+1)
				}
			}
			wsSubscribersMutex.Unlock()
			break

		default:
			broadcast(message, base.Type)
			break
		}
	}

	wsSubscribersMutex.Lock()
	delete(wsSubscribers, conn)
	wsSubscribersMutex.Unlock()
}

func wsServerStart() {
	http.HandleFunc("/chat", wsAccept)
	log.Fatal(http.ListenAndServe("0.0.0.0:5800", nil))
}
