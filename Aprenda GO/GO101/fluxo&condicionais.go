package main

import (
	"fmt"
	"log"
	"os"
)

var (
	cara, coroa int
)

func lancarMoeda(lado string) {
	switch lado {
	case "cara":
		cara++
	case "coroa":
		coroa++
	default:
		fmt.Println("Caiu em pé")
	}
}
func main() {
	a, b := 10, 10

	if a > b {
		println("A é maior do que B")
	} else if a < b {
		fmt.Println("A é menor do que B")
	} else {
		fmt.Println("A e B são iguais")
	}

	file, err := os.Open("hello.txt")
	if err != nil {
		log.Panic(err)
	}

	switch {
	case a > b:
		fmt.Println("A é maior que B")
	case b > a:
		fmt.Println("B é maior que A")
	default:
		fmt.Println("São iguais")
	}

	data := make([]byte, 100)

	if _, err := file.Read(data); err != nil {
		log.Panic(err)
	}

	fmt.Println(string(data))
}
