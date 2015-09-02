package main

import (
	"fmt"
	"math/rand"
)

var randData = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+{}[]|\\;:'<>,./?")

type Msg struct {
	buf       []byte
	prefixLen int
}

func NewMsg(size int) *Msg {
	prefix := fmt.Sprintf("rate=%v ", config.Rate)
	prefixLen := len(prefix)
	if prefixLen > size {
		prefixLen = size
	}

	msg := Msg{make([]byte, size), prefixLen}
	copy(msg.buf, prefix)

	return &msg
}

func (msg *Msg) Generate() string {
	for i := msg.prefixLen; i < len(msg.buf); i++ {
		msg.buf[i] = randData[rand.Intn(len(randData))]
	}
	return string(msg.buf)
}
