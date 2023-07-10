package main

import "fmt"

type Pessoa struct {
	Nome   string
	Idade  uint8
	Status bool
}

type PessoaFisica struct {
	Pessoa
	Sobrenome string
	cpf       string
}

func (p Pessoa) String() string {
	return fmt.Sprintf("Olá, meu nome é %s e eu tenho %d anos.", p.Nome, p.Idade)
}

type PessoaJuridica struct {
	RazaoSocial string
	cnpj        string
}

func main() {
	p := PessoaFisica{
		Pessoa{Nome: "Tiago", Idade: 31, Status: true},
		"temporin",
		"000.000.000.00",
	}
	fmt.Println(p)
}
