package main

import "fmt"

var (
	nome string
	N1   int
	n2   int32
)

func hello(nome string) {
	fmt.Println("Hello", nome, "!")
}

func sum(a, b int) int {
	return a + b
}

func main() {
	hello("Android")
	fmt.Println("Total:", sum(10, 5))
}
