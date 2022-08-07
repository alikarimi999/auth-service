package httpserver

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"time"
)

func generateId(n int) string {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letter[r.Intn(len(letter))]
	}
	return string(b)
}

func hashId(id string) string {
	h := sha256.New()
	h.Write([]byte(id))
	return hex.EncodeToString(h.Sum(nil))
}
