package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {

	// Numeros Inteiros
	fmt.Println("1, 2, 100")
	fmt.Println("Literal inteiro é", reflect.TypeOf(32000))

	// Sem sinal (só positivos)... uint8 uint6 uint32 uint64
	var b byte = 255
	fmt.Println("O byte é:", reflect.TypeOf(b))

	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é", i1)
	fmt.Println("O tipo do i1 é", reflect.TypeOf(i1))

	var i2 rune = 'a' //representa um mapeamentod da tabela Unicode (int32)
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println(i2)

	// Números reais (float32 float64)
	var x = float32(49.99)
	fmt.Println("O tipo de x é:", reflect.TypeOf(x))
	fmt.Println("O tipo literal 49.99 é", reflect.TypeOf(49.99)) //float64

	// Boolean
	bo := true
	fmt.Println("O tipo de bo é:", reflect.TypeOf(bo))
	fmt.Println(!bo)

	// String
	s1 := "Olá meu nome é vrau"
	fmt.Println(s1 + "!")
	fmt.Println("O tamanho da String é", len(s1))

	// String com multiplas linhas
	s2 := ` Olá 
	meu 
	nome 
	é
	Vrau`
	fmt.Println("O tamanho da string é", len(s2))

	// char???
	char := 'a'
	fmt.Println("O tipo de char é", reflect.TypeOf(char))
	fmt.Println(char)

}
