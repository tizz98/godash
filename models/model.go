package models

import (
	"math/rand"
	"time"
)

type Id string

type Model struct {
	Id Id `sql:"id"`
}

var src = rand.NewSource(time.Now().UnixNano())

const IdLength = 32
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// Generate a random id
// from: https://stackoverflow.com/a/31832326
func GenerateId() Id {
	b := make([]byte, IdLength)
	for i, cache, remain := IdLength-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return Id(b)
}
