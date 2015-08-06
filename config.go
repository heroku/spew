package main

import (
	"github.com/heroku/spew/Godeps/_workspace/src/github.com/kelseyhightower/envconfig"
	"log"
)

// Note: the `default` tag must appear before `envconfig` for the default thing
// to work.
type Config struct {
	MsgSize       int    `default:"32",envconfig:"MSGSIZE"`
	Seed          int64  `default:"8675309",envconfig:"SEED"`
	Rate          string `default:"1s",envconfig:"RATE"`
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
	// log.Printf("Config => %+v", config)
}
