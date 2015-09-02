package main

import (
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

func runOnQuit(fn func() int) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for _ = range c {
			// sig is a ^C, handle it
			os.Exit(fn())
		}
	}()
}

var randData = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+{}[]|\\;:'<>,./?")

func randStr(buf []byte) string {
	for i := 0; i < len(buf); i++ {
		buf[i] = randData[rand.Intn(len(randData))]
	}
	return string(buf)
}
