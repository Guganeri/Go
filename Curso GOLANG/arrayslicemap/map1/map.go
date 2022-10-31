package main

import "fmt"

func main() {
	//var aprovados map[int]string
	//mapas devem ser inicializados
	aprovados := make(map[int]string)

	aprovados[123456789] = "Maria"
	aprovados[121212122] = "Pedro"
	aprovados[123123123] = "Ana"

	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}
	fmt.Println(aprovados[123123123])
	delete(aprovados, 123123123)
	fmt.Println(aprovados[123123123])
}
