package twitch

func Start(out chan any) {
	go IrcPing()
	go Connect(out)
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
