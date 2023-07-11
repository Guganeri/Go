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

func (pf PessoaFisica) Doc() string {
	return fmt.Sprintf("me cpf é: %s", pf.cpf)
}

func (p Pessoa) String() string {
	return fmt.Sprintf("Olá, meu nome é %s e eu tenho %d anos.", p.Nome, p.Idade)
}

type PessoaJuridica struct {
	RazaoSocial string
	cnpj        string
}

type Document interface {
	Doc() string
}

func show(d Document) {
	fmt.Println(d)
	fmt.Println(d.Doc())
}

func main() {
	p := PessoaFisica{
		Pessoa{Nome: "Tiago", Idade: 31, Status: true},
		"temporin",
		"000.000.000.00",
	}
	fmt.Println(p)

	show(p)
}
