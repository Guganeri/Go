package main

type Car struct {
	Model string
	Color string
}

// metodo
func (c Car) Start() {
	println(c.Model + " has been is started")
}

func soma(x, y int) int {
	return x + y
}

func main() {
	car := Car{ //declarando e atribuindo a variável
		Model: "Ferrari",
		Color: "Red",
	}
	//car.Model = "Fiat"
	//println(car.Model)
	car.Start()
}
