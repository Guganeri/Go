package main

import (
	"fmt"
	"strings"
)

func SpinWords(input string) string {
	words := strings.Split(input, " ")
	fmt.Println(words)

	for i, word := range words {
		if len(word) >= 5 {
			words[i] = reverseWord(word)
		}
	}

	return strings.Join(words, " ")
}

func reverseWord(word string) string {

	length := len(word)
	arr := []byte(word)

	reversedWord := make([]byte, length)

	for i := 0; i < length; i++ {
		reversedWord[i] = arr[length-i-1]
	}

	wordFinal := string(reversedWord)

	return wordFinal

}

func main() {
	input := "Hey fellow warriors"
	reversed := SpinWords(input)
	fmt.Println(reversed)
}
