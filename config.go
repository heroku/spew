package main

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	MsgSize int    `envconfig:"MSGSIZE", default:32`
	Seed    int64  `envconfig:"SEED", default:8675309`
	Rate    string `envconfig:"RATE", default:"1s"`
}

func GetConfig() *Config {
	config := new(Config)
	err := envconfig.Process("spew", config)
	if err != nil {
		log.Fatalf("Incomplete config: %v", err)
	}
	return config
}
