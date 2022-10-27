package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	switch { //switch true, vai pegar o primeiro case verdadeiro
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde.")
	default:
		fmt.Println("Boa noite")
	}

}
