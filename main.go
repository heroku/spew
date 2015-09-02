package main

import (
	"fmt"
	"time"
)

func main() {
	runOnQuit(handleQuit)

	msg := NewMsg(config.MsgSize)
	for num := 1; ; num++ {
		time.Sleep(config.sleepTime)
		fmt.Println(num, "spews", msg.Generate())
	}
}
