package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"slices"
	"time"
)

func main() {
	go func() {
		for {
			fmt.Printf("%v\n", wsSubscribers)
			//for _, v := range wsSubscribers {
			//	fmt.Println(v)
			fmt.Println("----------")
			time.Sleep(time.Second * 10)
			//}
		}
	}()

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
