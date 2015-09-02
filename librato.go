package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func APIRequest(path string, jsonBody []byte) error {
	if config.LibratoUser == "" {
		return fmt.Errorf("Librato config not set")
	}
	url := "https://metrics-api.librato.com/v1" + path
	body := bytes.NewBuffer(jsonBody)
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return err
	}
	request.SetBasicAuth(config.LibratoUser, config.LibratoPass)
	request.Header.Set("Content-Type", "application/json")
	_, err = http.DefaultClient.Do(request)
	return err
}

type AnnotationEvent struct {
	Name        string `json:"-"`
	Title       string `json:"title"`
	Source      string `json:"source"`
	Description string `json:"description"`
	StartTime   int64  `json:"start_time"`
	EndTime     int64  `json:"end_time"`
}

func (evt *AnnotationEvent) Send() error {
	data, err := json.Marshal(evt)
	if err != nil {
		return err
	}
	return APIRequest("/annotations/"+evt.Name, data)
}

func Annotate(name, title, description string, start, end time.Time) error {
	evt := AnnotationEvent{
		Name:        name,
		Title:       title,
		Source:      config.LibratoSource,
		Description: description,
		StartTime:   start.Unix(),
		EndTime:     end.Unix(),
	}
	return evt.Send()
}

func handleQuit() int {
	if config.LibratoUser == "" {
		return 0
	}
	log.Printf("Received interrupt; sending annotation to librato...")
	description := fmt.Sprintf("Spew run with msg size = %d\nrate = %s\nseed = %d\nsource = %q",
		config.MsgSize, config.Rate, config.Seed, config.LibratoSource)
	endTime := time.Now()
	err := Annotate("spew-run", "Spew Run", description, startTime, endTime)
	if err != nil {
		log.Printf("ERROR sending to librato: %v", err)
		return 1
	} else {
		log.Printf("Done sending to librato.")
	}
	return 0
}

var startTime time.Time

func init() {
	startTime = time.Now()
}
