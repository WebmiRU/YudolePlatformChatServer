package goodgame

import (
	"YudolePlatofrmChatServer/obj"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"strconv"
	"strings"
)

var smiles = make(map[string]Smile)

func loadSmiles() {
	log.Println("SLS")
	smilesFile, err := os.ReadFile("./data/goodgame_smiles.json")

	if err != nil {
		log.Println("ERROR: Can't load GoodGame smile list")
	}

	var smilesData []SmileData

	err = json.Unmarshal(smilesFile, &smilesData)
	if err != nil {
		log.Println("ERROR: Can't parse file: \"goodgame_smiles.json\"")
	}

	for _, v := range smilesData {
		key := fmt.Sprintf(":%s:", v.Key)
		smiles[key] = Smile{
			Animated: v.Images.Gif,
			Static:   v.Images.Big,
		}
	}

	log.Println("SLS1")

	msg := processSmile("Hello :shoked: world")

	log.Println("SLS2")

	fmt.Println(msg)

	log.Println("SLS3")
}

func processSmile(message string) string {
	for k, v := range smiles {
		message = strings.ReplaceAll(message, k, fmt.Sprintf("<img src='%s' alt='%s' />", v, k))
	}

	return message
}

func connect(config obj.Config, out chan any) {
	ws, _, err := websocket.DefaultDialer.Dial("wss://chat-1.goodgame.ru/chat2/", nil)
	if err != nil {
		log.Println("GoodGame websocket server connection error", err)
	}

	defer ws.Close()

	for {
		var message Message
		err := ws.ReadJSON(&message)

		if err != nil {
			log.Println("WS read error:", err)
			return
		}

		switch message.Type {
		case "welcome":
			request := JoinRequest{
				Type: "join",
				Data: JoinRequestData{
					ChannelId: "9126",
					Hidden:    0,
					Mobile:    false,
					Reload:    false,
				},
			}

			ws.WriteJSON(&request)

			// Join to channels

			for _, v := range config.Services.GoodGame.Channels {
				request = JoinRequest{
					Type: "join",
					Data: JoinRequestData{
						ChannelId: strconv.FormatInt(v, 10),
						Hidden:    0,
						Mobile:    false,
						Reload:    false,
					},
				}

				ws.WriteJSON(&request)
			}

		case "success_join":
			fmt.Println("SUCCESS JOIN")

		case "channel_counters":
			fmt.Println("COUNTERS")

		case "message":
			msg := obj.ChatMessage{
				Type:    "chat/message",
				Service: "goodgame",
				Src:     message.Data.Text,
				Text:    message.Data.Text,
				Html:    message.Data.Text,
				User: obj.User{
					Name: message.Data.UserName,
					Meta: obj.UserMeta{
						Avatar: "",
						Badges: nil,
					},
				},
			}

			out <- msg

			log.Printf("[%s]: %s", message.Data.UserName, message.Data.Text)
		}
	}

}
