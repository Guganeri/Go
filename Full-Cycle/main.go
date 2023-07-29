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

	go worker(canal, 1)
	worker(canal, 2)

}

func worker(canal chan int, workerID int) {
	for {
		fmt.Println("Recebeu do canal", <-canal, "no worker", workerID)
		time.Sleep(time.Second)
	}
}
