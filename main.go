package main

import (
	"time"
	"fmt"
	"github.com/freeformz/shh/utils"
)


func main() {
	var duration = utils.GetEnvWithDefaultDuration("DURATION", "100ms")
	var i = 0

	ticker := time.Tick(duration)
	for t := range ticker {
		i++
		fmt.Println("Tick: ", i, " - ", t)
	}
}
