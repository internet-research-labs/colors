package main

import (
	"fmt"
	"github.com/carbocation/go-instagram/instagram"
	"image"
	_ "image/jpeg"
	"log"
	"net/http"
	// "os"
	"time"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetPrefix("[XXX] ")
}

// log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

// Wrapper for InstagramClient
type InstagramConnection struct {
	Client        instagram.Client
	Channel       chan image.Image
	Tag           string
	Started       bool
	Ticker        *time.Ticker
	MaxID         *string
	LastTimestamp *int64
}

// Create a New Instagram Connection
func NewInstagramConnection(client_id, tag string) *InstagramConnection {
	var returned InstagramConnection
	value := instagram.NewClient(nil)
	value.ClientID = client_id
	returned.Client = *value
	returned.Channel = make(chan image.Image)
	returned.Tag = tag
	returned.Started = false
	returned.Ticker = time.NewTicker(time.Second * 5)
	returned.MaxID = nil
	return &returned
}

func (s *InstagramConnection) Get() {
	var opt *instagram.Parameters

	// Set MaxID
	if s.LastTimestamp == nil {
		opt = &instagram.Parameters{Count: 25}
		log.Println("LastTimestamp is nil")
	} else {
		opt = &instagram.Parameters{Count: 25, MinTimestamp: *s.LastTimestamp}
		log.Println("LastTimestamp is", *s.LastTimestamp)
	}

	media, _, err := s.Client.Tags.RecentMedia(s.Tag, opt)

	if err != nil {
		fmt.Println(err)
		return
	}

	latest_timestamp := int64(-1)

	for _, val := range media {

		if val.CreatedTime > latest_timestamp {
			latest_timestamp = val.CreatedTime
		}

		url := val.Images.LowResolution.URL
		// fmt.Println(url)
		go func() {
			log.Printf("GET \"%s\"", url)
			s.Channel <- DownloadImage(url)
		}()
	}

	s.LastTimestamp = &latest_timestamp
}

func (s *InstagramConnection) Images() chan image.Image {
	return s.Channel
}

func DownloadImage(url string) image.Image {

	// Get body
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	m, _, err := image.Decode(resp.Body)
	defer func() {
		if recoverErr := recover(); err != nil {
			log.Println(recoverErr)
		}
	}()

	return m
}

// Start ticker
func (s *InstagramConnection) Start() {
	s.Started = true
	go s.StartTicking()
}

func (s *InstagramConnection) Stop() {
	s.Started = false
}

func (s *InstagramConnection) StartTicking() {
	for range s.Ticker.C {
		s.Get()
	}
}
