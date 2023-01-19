package main

import (
	"crypto/rand"
	"fmt"
)

func generateKey() {
	// Generate a 32-byte key
	key := make([]byte, 64)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println(err)
	}
	// Create a string with the characters that you want to include in the key
	characters := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+"
	keystring := ""
	// Iterate over the key slice to select a random character from the characters string
	for i := 0; i < 64; i++ {
		keystring += string(characters[key[i]%byte(len(characters))])
	}
	keystring += "Keep-calm-its-juts-a-test"
	fmt.Println(keystring)
}

func main() {
	generateKey()
}
