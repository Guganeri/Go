package main

import (
	"fmt"
	"time"
)

func processando() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

// T1
func main() {

	canal := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			canal <- 1
			fmt.Println("Jogou no canal", i)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			canal <- 1
			fmt.Println("Jogou no canal", i)
		}
	}()

	for x := range canal {
		fmt.Println("Recebeu no canal", x)
		time.Sleep(time.Second)
	}
}
