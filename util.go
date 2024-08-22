package main

import "math/rand"

// randomCode generates a random code of length n.
func randomCode(n int) string {
	// chars include base 58
	chars := "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	b := make([]byte, n)

	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}

	chars = string(b)

	return chars
}

func nonConflictCode(n int) string {
	// check db in the future
	return randomCode(n)
}
