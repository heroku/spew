package main

import (
	"time"
	"fmt"
	"github.com/freeformz/shh/utils"
)


func main() {
	var duration = utils.GetEnvWithDefaultDuration("DURATION", "100ms")
	ticker := time.Tick(duration)
	for t := range ticker {
		fmt.Println("Tick: ", t)
	}
}
