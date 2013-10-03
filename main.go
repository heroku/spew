package main

import (
	"time"
	"fmt"
	"github.com/freeformz/shh/utils"
)


func main() {
	var i = 0
	for {
		i++
		fmt.Println("Spew: ", i, " - ", time.Now())
	}
}
