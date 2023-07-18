package slices

import "fmt"

func main() {

	//nomes := make([]string, 10)

	idades := make(map[string]uint8)
	idades["Tiago"] = 31
	idades["Relnes"] = 22
	idades["Teste"] = 32

	val, ok := idades["Tiago"]
	fmt.Println(val, ok)

	//nomes := []string{"Tiago", "Dani", "Marcos", "Maria"}
	//fmt.Println(nomes, len(nomes), cap(nomes))

	//nomes = append(nomes, "Rafael")
	//fmt.Println(nomes, len(nomes), cap(nomes))

}
