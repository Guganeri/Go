package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{
		"José João":      1222.33,
		"Gabriela Silva": 123131.00,
		"Pedro Junior":   1231.1,
	}

	funcsESalarios["Rafael Filho"] = 1333.0
	delete(funcsESalarios, "Ineistente")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}

	for eduardo, monica := range funcsESalarios {
		fmt.Println(eduardo, monica)
	}

}
