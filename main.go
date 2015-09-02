package main

import (
	"fmt"
)

func main() {
	runOnQuit(handleQuit)

	msg := NewMsg(config.MsgSize)
	for {
		num := config.traffic.Next()
		fmt.Println(num, "spews", msg.Generate())
	}
}
