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
	go processando() // T2
	go processando() // T3
	processando()    // T1
}
