package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	defaultMsgSize = 32
	defaultSeed    = 8675309
	defaultRate    = 1 * time.Second
)

var bytes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+{}[]|\\;:'<>,./?")

func randStr(buf []byte) string {
	for i := 0; i < len(buf); i++ {
		buf[i] = bytes[rand.Intn(len(bytes))]
	}
	return string(buf)
}

func parseRate(rate string) (time.Duration, error) {
	var duration time.Duration
	var err error
	num := int64(1)
	bits := strings.Split(rate, "/")
	if len(bits) == 2 {
		duration, err = time.ParseDuration(bits[1])
		if err != nil {
			return defaultRate, err
		}
		num, err = strconv.ParseInt(bits[0], 10, 32)
		if err != nil {
			return defaultRate, err
		}
	} else {
		num = 1
		duration, err = time.ParseDuration(rate)
	}

	return time.Duration(int64(duration) / num), nil
}

func main() {
	var i = 0
	var buf []byte

	// Repeatable payloads.
	if seed, err := strconv.ParseInt(os.Getenv("SEED"), 10, 32); err == nil {
		rand.Seed(seed)
	} else {
		rand.Seed(defaultSeed)
	}

	sleepTime, err := parseRate(os.Getenv("RATE"))
	if err != nil {
		log.Println("Duration Parsing: ", err)
		log.Println("Continuing w/o duration")
	}

	mlen, err := strconv.ParseInt(os.Getenv("MSGSIZE"), 10, 32)
	if err == nil && mlen > 0 {
		buf = make([]byte, mlen)
	} else {
		log.Println("Unable to parse MSGSIZE, defaulting to MSGSIZE of", defaultMsgSize)
		buf = make([]byte, defaultMsgSize)
	}

	for {
		i++
		if sleepTime > 0 {
			time.Sleep(sleepTime)
		}

		fmt.Println(i, "spews", randStr(buf))
	}
}
