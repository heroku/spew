package main

import (
	"log"
	"os"
	"time"
)

func main() {
	var i = 0
	var duration time.Duration

	duration, err := time.ParseDuration(os.Getenv("DURATION"))
	if err != nil {
		log.Println("Duration Parsing: ", err)
		log.Println("Continueing w/o duration")
	}

	for {
		i++
		if duration > 0 {
			time.Sleep(duration)
		}
		log.Println("Spew: ", i)
	}
}
