package goodgame

import "YudolePlatofrmChatServer/obj"

func Start(config obj.Config, out chan any) {
	//go connect(config, out)
	loadSmiles()
}

func Reload() {

}

func Stop() {

}
