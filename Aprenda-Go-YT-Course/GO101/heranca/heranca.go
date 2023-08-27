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

func (pj PessoaJuridica) Doc() string {
	return fmt.Sprintf("me cpf é: %s", pj.cnpj)
}

func show(d Document) {
	switch d.(type) {
	case PessoaFisica:
		fmt.Println(d.(PessoaFisica).Sobrenome)
	case PessoaJuridica:
		fmt.Println(d.(PessoaJuridica).RazaoSocial)
	default:
		fmt.Println("tipo desconhecido")
	}

	fmt.Println(d)
	fmt.Println(d.Doc())
}

type Animal interface {
	Sound() string
	Eat(food string) string
}

type Dog struct {
	Name string
}

func (d Dog) Sound() string {
	return "Woof"
}

func (d Dog) Eat(food string) string {
	return fmt.Sprintf("%s is eating %s", d.Name, food)
}

func main() {
	p := PessoaFisica{
		Pessoa{Nome: "Tiago", Idade: 31, Status: true},
		"temporin",
		"000.000.000.00",
	}
	//fmt.Println(p)

	show(p)

	pj := PessoaJuridica{
		"Temporim LTDA",
		"123.123.123.12",
	}
	show(pj)

	//	var myAnimal Animal
	//	fmt.Println(myAnimal)
	//	myAnimal = Dog{Name: "Buddy"}
	//	fmt.Println(myAnimal.Sound())      // Output: Woof
	//	fmt.Println(myAnimal.Eat("bones")) // Output: Buddy is eating bones

}
