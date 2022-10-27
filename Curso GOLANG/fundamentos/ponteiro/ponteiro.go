package main

import "fmt"

func main() {
	i := 1

	var p *int = nil

	p = &i //pegando o edereco da variável
	*p++   // desreferenciando (pegando o valor)
	i++
	fmt.Println(p, *p, i, &i)

	// Go não tem aritmética de ponteiros
	//p++

}
