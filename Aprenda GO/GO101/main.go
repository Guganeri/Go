package main

import "fmt"

var (
	nome string
	N1   int
	n2   int32
)

func main() {

	var a, b, c = true, 2.3, "Olá"
	println(a, b, c)

	mensagem := "Aula 03"
	fmt.Println(mensagem)

	var total int
	total++
	fmt.Println("total:", total)
	nome = "Teste temp"
	fmt.Println("Hello", nome, "!")

	var x, y = 10, 20
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)
}
