package main

import (
	"YudoleChatServer/packages/resource"
	"crypto/sha256"
	"encoding/hex"
	"image"
	"io"
	"log"
	"net/http"
	"os"
)

type Resources struct {
	Audio  map[string]resource.Audio
	Images []resource.Image
}

func (r *Resources) Load() {
	r.LoadAudio()
	r.LoadImages()
}

func (r *Resources) LoadImages() {
	//if r.Images == nil {
	//	r.Images = make(map[string]resource.Image)
	//}

	for _, v := range config.Resources.Images {
		v.Source = "upload"
		v.Path = cd + ps + "data" + ps + "resources" + ps + "images" + ps + v.Sha256
		r.Images = append(r.Images, v)
	}

	for _, module := range modules {
		for _, v := range module.Resources.Images {
			path := module.Dir + "/resources/" + v
			res, err := r.GetImageFile(path)

			if err != nil {
				log.Println(err)
				continue
			}

			res.Source = "module"
			res.Path = path
			r.Images = append(r.Images, *res)
		}
	}
}

func (r *Resources) LoadAudio() {
	if r.Audio == nil {
		r.Audio = make(map[string]resource.Audio)
	}

	for _, v := range config.Resources.Audio {
		v.Source = "upload"
		v.Path = cd + ps + "data" + ps + "resources" + ps + "audio" + ps + v.Sha256
		r.Audio[v.Sha256] = v
	}

	for _, module := range modules {
		for _, v := range module.Resources.Audio {
			path := module.Dir + "/resources/" + v
			res, err := r.GetAudioFile(path)

			if err != nil {
				log.Println(err)
				continue
			}

			res.Source = "module"
			res.Path = path
			r.Audio[res.Sha256] = *res
		}
	}
}

func (r *Resources) GetAudioFile(path string) (*resource.Audio, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	s256 := sha256.New()
	io.Copy(s256, file)
	sha256String := hex.EncodeToString(s256.Sum(nil))

	stat, _ := file.Stat()
	buff := make([]byte, 100)
	file.ReadAt(buff, 0)
	mimeType := http.DetectContentType(buff)

	result := resource.Audio{
		Name:     stat.Name(),
		Size:     stat.Size(),
		Sha256:   sha256String,
		MimeType: mimeType,
	}

	file.Close()

	return &result, nil
}

func (r *Resources) GetImageFile(path string) (*resource.Image, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	s256 := sha256.New()
	io.Copy(s256, file)
	sha256String := hex.EncodeToString(s256.Sum(nil))

	stat, _ := file.Stat()
	buff := make([]byte, 100)
	file.ReadAt(buff, 0)

	name := stat.Name()
	size := stat.Size()
	mimeType := http.DetectContentType(buff)

	file.Close()

	imageFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	im, _, err := image.DecodeConfig(imageFile)

	width := im.Width
	height := im.Height

	imageFile.Close()

	result := resource.Image{
		Name:     name,
		Size:     size,
		Sha256:   sha256String,
		MimeType: mimeType,
		Width:    width,
		Height:   height,
	}

	return &result, nil
}
