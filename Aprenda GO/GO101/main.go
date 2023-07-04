package main

import (
	"fmt"
	"strconv"
)

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

func convertAndSum(a int, b string) (total int, err error) {
	i, err := strconv.Atoi(b)
	if err != nil {
		fmt.Println("ERROR", err)
		return
	}

	total = a + i
	return

}

func main() {
	hello("Android")
	fmt.Println("Total:", sum(10, 5))
	total, err := convertAndSum(10, "ia")
	if err != nil {
		fmt.Println("Utilize um valor válido")
	}
	fmt.Println("total:", total)
}
