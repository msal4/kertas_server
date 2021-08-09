package util

import "math/rand"

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandomString generates a random string of length n.
func RandomString(source rand.Source, n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[source.Int63()%int64(len(letters))]
	}
	return string(b)
}
