package metodos

import "fmt"

func main() {

	type Pessoa struct {
		Nome 	  string
		Sobrenome string
		Idade	  uint8
		Status    bool
		cpf       string
	}

	func (p Pessoa) String() string {
		return fmt.Sprintf("Olá meu nome é %s e eu tenho %d anos.", p.Nome, p.Idade)
	}


	type Categoria struct {
		Nome string
		Pai *Categoria
	}

	func (c Categoria) HasParent() bool {
		return c.Pai != nil
	}

	func main(){
		p := Pessoa{"Tiago", "Temporin", 31, true, "000.000.000.000"}
		p.cpf ="1"


	}
}
