package main

import (
	"fmt"
)

func main() {
	runOnQuit(handleQuit)

	msg := NewMsg(config.MsgSize)
	for {
		num, msgPrefix := config.traffic.Next()
		fmt.Println(num, "spews", msgPrefix+msg.Generate())
	}
}
