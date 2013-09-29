package main

import (
	"time"
	"fmt"
)

func main() {
	ticker := time.Tick(100 * time.Millisecond)
	for t := range ticker {
		fmt.Println("Tick: ", t)
	}
}
