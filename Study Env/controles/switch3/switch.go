package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "inteiro"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case func():
		return "funçao"
	default:
		return "não sei"
	}
}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(1))
	fmt.Println(tipo("opa"))
	fmt.Println(tipo(func() {}))
	fmt.Print(tipo(time.Nanosecond))
}
