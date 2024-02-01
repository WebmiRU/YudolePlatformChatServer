package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"yudole.ru/chatsdk/yclient"
	"yudole.ru/chatsdk/yconfig"
)

func main() {
	yconfig.Load()
	go yclient.Connect()
	wsc()
}

func wsc() {
	ws, _, err := websocket.DefaultDialer.Dial("wss://chat-1.goodgame.ru/chat2/", nil)
	if err != nil {
		log.Fatal("GoodGame websocket server connection error", err)
	}

	defer ws.Close()

	for {
		var message GoodgameMessage
		err := ws.ReadJSON(&message)

		if err != nil {
			log.Println("WS read error:", err)
			return
		}

		switch message.Type {
		case "welcome":
			request := GoodgameJoinRequest{
				Type: "join",
				Data: GoodgameJoinRequestData{
					ChannelId: "9126",
					Hidden:    0,
					Mobile:    false,
					Reload:    false,
				},
			}

			ws.WriteJSON(&request)

			request = GoodgameJoinRequest{
				Type: "join",
				Data: GoodgameJoinRequestData{
					ChannelId: "53029",
					Hidden:    0,
					Mobile:    false,
					Reload:    false,
				},
			}

			ws.WriteJSON(&request)

		case "success_join":
			fmt.Println("SUCCESS JOIN")

		case "channel_counters":
			fmt.Println("COUNTERS")

		case "message":
			yclient.Out <- yclient.Message{
				Id:      "",
				Type:    "chat/message",
				Service: "goodgame",
				Html:    message.Data.Text,
				Text:    message.Data.Text,
				User: yclient.User{
					Id:       "",
					Nickname: message.Data.UserName,
					Login:    message.Data.UserName,
					Meta: yclient.Meta{
						Badges: nil,
					},
				},
			}
			log.Printf("[%s]: %s", message.Data.UserName, message.Data.Text)
		}
	}

}
