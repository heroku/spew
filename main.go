package main

import (
	"time"
	"fmt"
)


func main() {
	var i = 0
	for {
		i++
		fmt.Println("Spew: ", i, " - ", time.Now())
	}
}
