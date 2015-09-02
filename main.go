package main

import (
	"fmt"
	"time"
)

func main() {
	runOnQuit(handleQuit)

	buf := make([]byte, config.MsgSize)
	for num := 1; ; num++ {
		time.Sleep(config.sleepTime)
		fmt.Println(num, "spews", randStr(buf))
	}
}
