package main

import (
	"YudolePlatofrmChatServer/obj"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

var config obj.Config
var out = make(chan any, 9999)

func configLoad() {
	file, err := os.ReadFile("config.json")

	if err != nil {

	}

	if err := json.Unmarshal(file, &config); err != nil {
		return
	}

}

func configSave() {
	cfg, _ := json.Marshal(config)
	if err := os.WriteFile("config.json", cfg, 0644); err != nil {
		// @TODO Error
		return
	}
}

func main() {
	configLoad()

	//go test1()
	//go test2()
	//go test3()
	//go test4()
	//go test5()
	//go test6()

	go wsServerStart()
	go twitch.Start(config, out)

	broadcast()
	//select {}
}

func test1() {
	//client := types.Client{
	//	Write: func(message json.RawMessage) error {
	//		fmt.Println("MSG!:", string(message))
	//
	//		return nil
	//	},
	//}
	//
	//Clients = append(Clients, &client)

	for {
		n := 0

		for _, client := range obj.Clients {
			n = n + 1
			m := fmt.Sprintf("Message 00%d", n)
			var x json.RawMessage
			x, _ = json.Marshal(m)

			var f = client.Write

			if err := f(x); err != nil {
				// Error
			}
		}

		time.Sleep(1 * time.Second)
	}
}
func test2() {
	for {
		s := obj.ChatMessage{
			Type:    "1",
			Service: "2",
			Src:     "3",
			Text:    "4",
			Html:    "5",
			User:    obj.User{},
		}
		//m, _ := json.Marshal(s)
		out <- s

		time.Sleep(2 * time.Second)
	}
}
func test3() {
	for {
		fmt.Println(len(out))
		time.Sleep(1 * time.Second)
	}
}
func test4() {
	for {
		var msg = <-out
		fmt.Println(msg)
	}
}
func test5() {
	for {
		fmt.Printf("CLIENTS Count: %d\n", len(obj.Clients))
		time.Sleep(1 * time.Second)
	}
}
func test6() {
	for {
		time.Sleep(time.Second * 3)
		fmt.Println(config)
	}
}

func broadcast() {
	for {
		m := <-out
		message, _ := json.Marshal(m)
		fmt.Println("OUT:", string(message))

		for _, client := range obj.Clients {
			if err := client.Write(message); err != nil {
				// Error
			}
		}
	}
}
