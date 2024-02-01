package goodgame

import "YudolePlatofrmChatServer/obj"

func Start(config obj.Config, out chan any) {
	loadSmiles()
	go connect(config, out)
	//loadSmiles()
}

func Reload() {

}

func Stop() {

}
