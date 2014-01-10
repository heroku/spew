package main

import (
	"log"
	"log/syslog"
	"os"
	"time"
)

func main() {
	var i = 0
	var duration time.Duration
	var sl *log.Logger
	longString := "0123456789"

	long := os.Getenv("LONG") != ""

	if long {
		for i := 0; i < 12; i++ {
			longString = longString + "-" + longString
		}
		longString = longString + "END"
	}

	log.SetFlags(log.Ldate | log.Lmicroseconds)

	duration, err := time.ParseDuration(os.Getenv("DURATION"))
	if err != nil {
		log.Println("Duration Parsing: ", err)
		log.Println("Continuing w/o duration")
	}

	if os.Getenv("SYSLOG") != "" {
		sl, err = syslog.NewLogger(syslog.LOG_ALERT|syslog.LOG_USER, log.Ldate|log.Lmicroseconds)
		if err != nil {
			log.Println("Error setting up syslog: ", err)
			log.Println("Continuing w/o syslog")
		}
	}

	for {
		i++
		if duration > 0 {
			time.Sleep(duration)
		}
		if long {
			log.Println("Spew: ", i, len(longString), longString)
		} else {
			log.Println("Spew: ", i)
		}
		if sl != nil {
			sl.Println("Syslog Spew: ", i)
		}
	}
}
