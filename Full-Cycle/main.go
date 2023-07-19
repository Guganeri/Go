package main

import "github.com/guganeri/Go/Full-Cycle/internal/entity"

type Car struct {
	Model string
	Color string
}

// metodo
func (c Car) Start() {
	println(c.Model + " has been is started")
}

func (c *Car) ChangeColor(color string) {
	c.Color = color
	println("New color: " + c.Color)
}

func main() {

	order := entity.NewOrder("1", 10, 1)

	car := Car{ //declarando e atribuindo a variável
		Model: "Ferrari",
		Color: "Red",
	}
	//car.Model = "Fiat"
	//println(car.Model)
	car.Start()
	car.ChangeColor("Blue")

	println(car.Color)
}
