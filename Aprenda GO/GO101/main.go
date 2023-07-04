package main

import "fmt"

var (
	nome string
	N1   int
	n2   int32
)

func main() {

	mensagem := "Aula 03"
	fmt.Println(mensagem)

	var total int
	total++
	fmt.Println("total:", total)
	nome = "Teste temp"
	fmt.Println("Hello", nome, "!")
}
