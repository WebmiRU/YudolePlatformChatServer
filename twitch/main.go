package twitch

import "YudolePlatofrmChatServer/obj"

func Start(config obj.Config, out chan any) {
	go IrcPing()
	go Connect(config, out)
	//fmt.Println("CONNECTED!")

	//for {
	//	var msg = <-*out
	//	message, _ := json.Marshal(msg)
	//	fmt.Println("OUT:", string(message))
	//}
}

func Reload() {

}

func Stop() {

}
