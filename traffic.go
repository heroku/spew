package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Traffic struct {
	normalSleep time.Duration // Sleep time between normal log records
	spike       struct {
		enabled bool
		sleep   time.Duration // Sleep time between consequetive spikes
		length  int           // Length of each spike
	}
	counter int
}

func NewTraffic(rate string, spikeSpec string) (*Traffic, error) {
	var err error
	var traffic Traffic

	traffic.normalSleep, err = parseRate(rate)
	if err != nil {
		return nil, err
	}

	if spikeSpec == "" {
		traffic.spike.enabled = false
	} else {
		traffic.spike.enabled = true
		traffic.spike.sleep, traffic.spike.length, err = parseSpikeSpec(spikeSpec)
		if err != nil {
			return nil, fmt.Errorf("Invalid spike spec (%v), must be format: <waittime>:<length>", err)
		}
	}

	return &traffic, nil
}

func (t *Traffic) Next() int {
	time.Sleep(t.normalSleep)
	t.counter += 1
	return t.counter
}

// parseRate converts a rate string like "100/1s" to "10ms" duration value
func parseRate(rate string) (time.Duration, error) {
	var duration time.Duration
	var err error
	num := float64(1)
	bits := strings.Split(rate, "/")
	if len(bits) == 2 {
		duration, err = time.ParseDuration(bits[1])
		if err != nil {
			return 0, err
		}
		num, err = strconv.ParseFloat(bits[0], 32)
		if err != nil {
			return 0, err
		}
	} else {
		num = 1
		duration, err = time.ParseDuration(rate)
		if err != nil {
			return 0, err
		}
	}

	return time.Duration(float64(duration) / float64(num)), err
}

// Example: "10s:300"  -> (10s, 300, nil)  i.e., spike of 300 every 10s
func parseSpikeSpec(spikeSpec string) (time.Duration, int, error) {
	bits := strings.SplitN(spikeSpec, ":", 2)
	if len(bits) < 2 {
		return 0, 0, fmt.Errorf("need two-part")
	}

	duration, err := time.ParseDuration(bits[0])
	if err != nil {
		return 0, 0, err
	}

	length, err := strconv.Atoi(bits[1])
	if err != nil {
		return 0, 0, err
	}

	return duration, length, err
}
