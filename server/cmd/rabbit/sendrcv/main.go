package main

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	c, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	q, err := c.QueueDeclare(
		"go_q1",
		true,  // durable
		false, //autoDelete
		false, // exclusive
		false, //nowait
		nil,   // args
	)
	if err != nil {
		panic(err)
	}
	go consume("c1", conn, q.Name)
	go consume("c2", conn, q.Name)
	i := 0
	for {
		i++
		err = c.Publish(
			"",
			q.Name,
			false,
			false,
			amqp.Publishing{
				Body: []byte(fmt.Sprintf("message %d", i)),
			},
		)
		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(time.Millisecond * 200)
	}
}

func consume(consume string, conn *amqp.Connection, q string) {
	c, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	msgs, err := c.Consume(
		q,
		consume,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}
	for msg := range msgs {
		fmt.Printf("%s: %s\n", consume, msg.Body)
	}
}
