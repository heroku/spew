package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var randData = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+{}[]|\\;:'<>,./?")

func generateRandString(n int) string {
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = randData[rand.Intn(len(randData))]
	}
	buf[0] = 'M'
	return string(buf)
}

type Msg struct {
	size    int
	metrics []string
}

func NewMsg(size int) *Msg {
	return &Msg{
		size,
		[]string{fmt.Sprintf("rate=%v", config.Rate)},
	}
}

func (msg *Msg) Generate(extraMetrics ...string) string {
	metrics := append(msg.metrics, extraMetrics...)
	prefix := strings.Join(metrics, " ") + " "
	if len(prefix) > msg.size {
		return prefix[:msg.size]
	} else {
		return prefix + generateRandString(msg.size-len(prefix))
	}
}
