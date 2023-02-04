package main

import "fmt"

func inc1(n int) {
	n++
}

// revisão: um ponteiro é representado por um *
func inc2(n *int) {
	//revisão: * é usado para desferenciar, ou seja ter acesso ao valor o qual o ponteiro aponta
	*n++
}

func main() {
	n := 1
	inc1(n) //por valor

	fmt.Println(n)

	//revisão: & usado para obter o endereço da variável
	inc2(&n)
	fmt.Println(n)

}
