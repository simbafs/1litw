package main

import (
	"math/rand"
	"net/url"
)

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

// IsValidURL checks if a string is a valid URL and the scheme is http or https.
func IsValidURL(str string) bool {
	u, err := url.Parse(str)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return u.Scheme == "http" || u.Scheme == "https"
}
