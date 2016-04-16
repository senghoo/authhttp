package main

import (
	"math/rand"
	"time"
)

const seedOffset = 34052689

func init() {
	rand.Seed(time.Now().UnixNano() & seedOffset)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
