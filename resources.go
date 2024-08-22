package main

import "YudoleChatServer/packages/resource"

type Resources struct {
	Audio  map[string]resource.Audio
	Images map[string]resource.Image
}

func (r *Resources) Load() {
	r.LoadImages()
	r.LoadAudio()
}

func (r *Resources) LoadImages() {

}

func (r *Resources) LoadAudio() {

}
