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
	c.ExchangeDeclare(
		"go_ex",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	go subscrib(conn, "go_ex")
	go subscrib(conn, "go_ex")
	i := 0
	for {
		i++
		err = c.Publish(
			"go_ex",
			"", //routing key
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
func subscrib(conn *amqp.Connection, ex string) {
	c, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	q, err := c.QueueDeclare(
		"",
		false, // durable
		true,  //autoDelete
		false, // exclusive
		false, //nowait
		nil,   // args
	)
	if err != nil {
		panic(err)
	}
	defer c.QueueDelete(
		q.Name,
		false,
		false,
		false,
	)
	c.QueueBind(
		q.Name,
		"",
		"go_ex",
		false,
		nil,
	)
	consume("c", c, q.Name)
}
func consume(consume string, c *amqp.Channel, q string) {
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
