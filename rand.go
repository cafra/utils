package utils

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaNum = []byte(`0123456789`)
var alphCardNum = []byte(`1A2B3C4D5E6F7G8W9X1Y2Z3H4J5K6L7M8N9P1Q2I3R4S5T6U7V`)

func RandomCreateBytes(n int, alphabets ...byte) string {
	if len(alphabets) == 0 {
		alphabets = alphaNum
	}
	var bytes = make([]byte, n)
	var randBy bool
	if num, err := rand.Read(bytes); num != n || err != nil {
		r.Seed(time.Now().UnixNano())
		randBy = true
	}
	for i, b := range bytes {
		if randBy {
			bytes[i] = alphabets[r.Intn(len(alphabets))]
		} else {
			bytes[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return B2S(bytes)
}
