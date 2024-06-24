package main

import (
	"YudoleChatServer/packages/resource"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"slices"
	"sync"
)

type sseClient struct {
	W       *http.ResponseWriter
	R       *http.Request
	Chan    *chan any
	Channel string
	Events  []string
}

func (c *sseClient) Send(message any) error {
	*c.Chan <- message
	return nil
}

func (c *sseClient) Drop() error {
	sseClientsMutex.Lock()
	idx := slices.Index(sseClients, c)
	sseClients = slices.Delete(sseClients, idx, idx+1)
	sseClientsMutex.Unlock()

	//sseEventSubsMutex.Lock()
	//for event, clients := range sseEventSubs {
	//	idx := slices.Index(clients, c)
	//
	//	if idx == -1 {
	//		continue
	//	}
	//
	//	sseEventSubs[event] = slices.Delete(clients, idx, idx+1)
	//}
	//sseEventSubsMutex.Unlock()

	return nil
}

// SendConfig Функция отправки настроек канала SSE клиенту
func (c *sseClient) SendConfig() error {
	if _, ok := config.Channels[c.Channel]; ok {
		c.Send(resource.ChannelConfig{
			Type:    "system/channel/config",
			Payload: config.Channels[c.Channel],
		})
	}

	return nil
}

var sseClientsMutex sync.Mutex
var sseClients = make([]*sseClient, 0)
var sseChan = make(chan Message)

func sseChannelGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Content-Type", "text/event-stream")

	vars := mux.Vars(r)
	channel := vars["channel"]
	ch := make(chan any)
	var events []string

	if channel == "system" {
		events = []string{
			"event/unsubscribe",
			"stream/chat/private_message",
			"api/modules/update/index",
		}
	} else {
		events = config.Channels[channel].Events
	}

	client := &sseClient{
		W:       &w,
		R:       r,
		Chan:    &ch,
		Channel: channel,
		Events:  events,
	}

	sseClientsMutex.Lock()
	sseClients = append(sseClients, client)
	sseClientsMutex.Unlock()

	go func() {
		if err := client.SendConfig(); err != nil {
			log.Println(err)
			return
		}
	}()

loop:
	for {
		select {
		case message := <-ch:
			msg, _ := json.Marshal(message)
			if _, err := fmt.Fprintf(w, "data: %s\n\n", msg); err != nil {
				break loop
			}

			w.(http.Flusher).Flush()

		case <-r.Context().Done():
			break loop
		}
	}

	client.Drop()
	fmt.Println("SSE CLIENT DISCONNECTED")
}

func apiEventsHandlerGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Cache-Control", "no-cache")

	msg, _ := json.Marshal(resource.EventIndex{Payload: events})
	w.Write(msg)

	w.(http.Flusher).Flush()
}
