package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     uint8
	Status    bool
	cpf       string
}

func (p Pessoa) String() string {
	return fmt.Sprintf("Olá, %s! Você tem %d anos.", p.Nome, p.Idade)
}

type Categoria struct {
	Nome string
	Pai  *Categoria
}

func (c Categoria) HasParent() bool {
	return c.Pai != nil
}

func (c *Categoria) MeuMetodo() {

}

func main() {
	p := Pessoa{"Tiago", "Temporin", 31, true, "000.000.000.000"}
	p.cpf = "1"
	fmt.Println(p)

	cat := Categoria{Nome: "Categoria 1", Pai: &Categoria{Nome: "Pai"}}
	if !cat.HasParent() {
		fmt.Println("Não tem pai")
	}
}
