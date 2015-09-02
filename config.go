package main

import (
	"github.com/heroku/spew/Godeps/_workspace/src/github.com/kelseyhightower/envconfig"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// Note: the `default` tag must appear before `envconfig` for the default thing
// to work.
type Config struct {
	MsgSize       int    `default:"32" envconfig:"MSGSIZE"`
	Seed          int64  `default:"8675309" envconfig:"SEED"`
	Rate          string `default:"1s" envconfig:"RATE"`
	sleepTime     time.Duration
	LibratoUser   string `envconfig:"LIBRATO_USER"`
	LibratoPass   string `envconfig:"LIBRATO_PASS"`
	LibratoSource string `envconfig:"LIBRATO_SOURCE"`
}

var config Config

func init() {
	err := envconfig.Process("spew", &config)
	if err != nil {
		log.Fatalf("Incomplete config: %v", err)
	}
	config.sleepTime, err = parseRate(config.Rate)
	if err != nil {
		log.Fatalf("Invalid value for RATE (%v): %v", config.Rate, err)
	}
	rand.Seed(config.Seed)
}

// parseRate converts a rate string like "100/1s" to "10ms" duration value
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
