package main

import "fmt"

type nota float64

func (n nota) entre(inicio, fim float64) bool {
	return float64(n) >= inicio && float64(n) <= fim

}

func notaParaConceito(n nota) string {
	if n.entre(9.0, 10.0) {
		return "A"
	} else if n.entre(7.0, 8.0) {
		return "B"
	} else if n.entre(6.0, 7.0) {
		return "C"
	} else {
		return "E"
	}

}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(7.1))
	fmt.Println(notaParaConceito(6.0))
	fmt.Println(notaParaConceito(3.0))
}
