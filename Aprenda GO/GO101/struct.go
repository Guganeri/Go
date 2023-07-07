package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     uint8
	Status    bool
	cpf       string
}

func main() {
	p := Pessoa{
		Nome:      "Tiago",
		Sobrenome: "Temporim",
		Idade:     31,
		Status:    true,
		cpf:       "300.300.300.33",
	}

	fmt.Println(p, p.Nome, p.Idade, p.cpf)

}
