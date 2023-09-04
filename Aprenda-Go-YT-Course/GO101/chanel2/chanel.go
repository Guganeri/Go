package main

import (
	"fmt"
	"time"
)

func numeros(n chan<- int) {
	for i := 0; i < 10; i++ {
		n <- i
		fmt.Printf("Escrito no chanel: %d\n", i)
		time.Sleep(time.Millisecond * 150)
	}
	close(n)
}

func main() {
	cn := make(chan int)
	go numeros(cn)

	for v := range cn {
		fmt.Printf("lido do chanel: %d\n", v)
		time.Sleep(time.Millisecond * 130)
	}
	fmt.Println("Fim da execução")
}
