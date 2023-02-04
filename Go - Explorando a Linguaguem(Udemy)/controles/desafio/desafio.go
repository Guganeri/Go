package main

import "fmt"

func notaParaConceito(n float64) string {
	switch {

	case n > 10:
		return "Nota inválida"
	case n == 10 || n >= 9.9:
		return "A"
	case n == 8 || n >= 7:
		return "B"
	case n == 6 || n >= 5:
		return "C"
	default:
		return "E"
	}
}

func main() {

	fmt.Println(notaParaConceito(3))
}
