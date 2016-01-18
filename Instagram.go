package main

import (
	"fmt"
	"github.com/carbocation/go-instagram/instagram"
	"image"
	"net/http"
)

var URL string = "https://api.instagram.com/v1/tags/%s/media/recent?client_id=%s"

type InstagramConnection struct {
	Client  instagram.Client
	Channel chan image.Image
	Tag     string
}

func NewInstagramConnection(client_id, tag string) *InstagramConnection {
	var returned InstagramConnection
	value := instagram.NewClient(nil)
	value.ClientID = client_id
	returned.Client = *value
	returned.Channel = make(chan image.Image)
	returned.Tag = tag
	return &returned
}

func (s *InstagramConnection) Get() {
	opt := &instagram.Parameters{Count: 25}
	media, _, _ := s.Client.Tags.RecentMedia(s.Tag, opt)
	fmt.Println(media)
	fmt.Println(media[0].Images.LowResolution.URL)
	fmt.Println(media[1].Images.LowResolution.URL)
	for _, val := range media {
		url := val.Images.LowResolution.URL
		go func() {
			s.Channel <- DownloadImage(url)
		}()
	}
	fmt.Println("---")
}

func (s *InstagramConnection) Images() chan image.Image {
	return s.Channel
}

func DownloadImage(url string) image.Image {
	resp, _ := http.Get(url)
	m, _, _ := image.Decode(resp.Body)
	defer resp.Body.Close()
	return m
}
