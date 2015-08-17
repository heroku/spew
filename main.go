package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

var randData = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+{}[]|\\;:'<>,./?")

func randStr(buf []byte) string {
	for i := 0; i < len(buf); i++ {
		buf[i] = randData[rand.Intn(len(randData))]
	}
	return string(buf)
}

func parseRate(rate string) (time.Duration, error) {
	var duration time.Duration
	var err error
	num := float64(1)
	bits := strings.Split(rate, "/")
	if len(bits) == 2 {
		duration, err = time.ParseDuration(bits[1])
		if err != nil {
			return 0, err
		}
		num, err = strconv.ParseFloat(bits[0], 32)
		if err != nil {
			return 0, err
		}
	} else {
		num = 1
		duration, err = time.ParseDuration(rate)
		if err != nil {
			return 0, err
		}
	}

	return time.Duration(float64(duration) / float64(num)), err
}

func main() {
	var i = 0
	var buf []byte
	startTime := time.Now()
	description := fmt.Sprintf("Spew run with msg size = %d\nrate = %s\nseed = %d\nsource = %q",
		config.MsgSize, config.Rate, config.Seed, config.LibratoSource)

	// Signal handling
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for _ = range c {
			// sig is a ^C, handle it
			log.Printf("Received interrupt; sending annotation to librato...")
			endTime := time.Now()
			err := Annotate("spew-run", "Spew Run", description, startTime, endTime)
			if err != nil {
				log.Printf("ERROR sending to librato: %v", err)
				os.Exit(1)
			} else {
				log.Printf("Done sending to librato.")
				os.Exit(0)
			}
		}
	}()

	// Repeatable payloads.
	rand.Seed(config.Seed)

	sleepTime, err := parseRate(config.Rate)
	if err != nil {
		log.Fatalf("Invalid value for RATE (%v): %v", config.Rate, err)
	}

	buf = make([]byte, config.MsgSize)

	for {
		i++
		if sleepTime > 0 {
			time.Sleep(sleepTime)
		}

		fmt.Println(i, "spews", randStr(buf))
	}
}
