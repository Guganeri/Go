package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print("Linha. ")

	fmt.Println("Nova")
	fmt.Println("Linha.")

	x := 3.12121312

	//fmt.Print("O valor de  é" + x)

	xs := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xs)

	fmt.Println("O valor de x é", x)

	fmt.Printf("O valor ed x é %.2f .", x)

	a := 1
	b := 1.9999
	c := false
	d := "opa"
	fmt.Printf("\n%d %f %.1f %t %s", a, b, b, c, d)
	fmt.Printf("\n%v %v %v %v", a, b, c, d)
}
