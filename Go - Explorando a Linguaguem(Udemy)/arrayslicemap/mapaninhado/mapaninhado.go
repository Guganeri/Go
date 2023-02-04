package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15333.21,
			"Guga Pareira":   12112.12,
		},
		"J": {
			"José João": 111221.54,
		},
		"P": {
			"Pedro Junior": 22222.11,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
