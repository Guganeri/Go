package main

import (
	"fmt"
	"time"
)

func numeros() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
		time.Sleep(time.Millisecond * 150)
	}
}

func main() {
	numeros()

}
