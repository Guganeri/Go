package main

import "fmt"

func main() {
	nomes := []string{"Tiago", "Daniel", "Maria", "Marta"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	var i int

	//	for i < len(nomes) {
	//		fmt.Println(nomes[i])
	//		i++
	//	}

	for {
		fmt.Println(nomes[i])
	}
	//	for i := 0; i < len(nomes); i++ {
	//		fmt.Println(nomes[i])
	//	}

	//	for i := 0; i < 10; i++ {
	//		fmt.Println(i)
	//	}
}
