package main

import (
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
