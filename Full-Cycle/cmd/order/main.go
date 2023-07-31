package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/guganeri/Go/Full-Cycle/internal/infra/database"
	"github.com/guganeri/Go/Full-Cycle/internal/usecase"
	"github.com/guganeri/Go/Full-Cycle/pkg/rabbitmq"
	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()
	msgRabbitmqChannel := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgRabbitmqChannel) // escutando a fila //trava

	input := usecase.OrderInput{
		ID:    "123",
		Price: 10.0,
		Tax:   1.0,
	}
	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)

	rabbitmqWorder(msgRabbitmqChannel, uc)
}

func rabbitmqWorder(msgChan chan amqp.Delivery, uc *usecase.CalculateFinalPrice) {
	fmt.Println("Starting rabbitmq")
	for msg := range msgChan {
		var input usecase.OrderInput
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			panic(err)
		}
		output, err := uc.Execute(input)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)
		fmt.Println("Msg processado e salva no banco: ", output)
	}
}
