package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	MsgSize int    `envconfig:"MSGSIZE", default:32`
	Seed    int64  `envconfig:"SEED", default:8675309`
	Rate    string `envconfig:"RATE", default:"1s"`
}

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
			return 0, err
		}
		num, err = strconv.ParseInt(bits[0], 10, 32)
		if err != nil {
			return 0, err
		}
	} else {
		num = 1
		duration, err = time.ParseDuration(rate)
	}

	return time.Duration(int64(duration) / num), err
}

func main() {
	var i = 0
	var buf []byte
	var config Config

	err := envconfig.Process("spew", &config)
	if err != nil {
		log.Fatalf("Environment config error: %v", err)
	}

	// Repeatable payloads.
	rand.Seed(config.Seed)

	sleepTime, err := parseRate(config.Rate)
	if err != nil {
		log.Fatalf("Invalid value for DURATION: %v", err)
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
