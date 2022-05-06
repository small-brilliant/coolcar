package amqpclient

import (
	"context"
	carpb "coolcar/car/api/gen/v1"
	"encoding/json"
	"fmt"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
)

type Publisher struct {
	ch       *amqp.Channel
	exchange string
}

func NewPublisher(conn *amqp.Connection, ex string) (*Publisher, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("cannot allocate channel: %v", err)
	}
	err = declareExchange(ch, ex)
	if err != nil {
		return nil, fmt.Errorf("cannot declare exchange: %v", err)
	}
	return &Publisher{
		ch:       ch,
		exchange: ex,
	}, nil
}

func (p *Publisher) Publish(c context.Context, car *carpb.CarEntity) error {
	b, err := json.Marshal(car)
	if err != nil {
		return fmt.Errorf("cannot marshal: %v", err)
	}
	return p.ch.Publish(
		p.exchange,
		"",
		false,
		false,
		amqp.Publishing{
			Body: b,
		},
	)
}

type Subscriber struct {
	conn     *amqp.Connection
	exchange string
	logger   *zap.Logger
}

func NewSubscriber(conn *amqp.Connection, exchange string, logger *zap.Logger) (*Subscriber, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("cannot allocate channel: %v", err)
	}
	defer ch.Close()
	err = declareExchange(ch, exchange)
	if err != nil {
		return nil, fmt.Errorf("cannot declare exchange: %v", err)
	}
	return &Subscriber{
		conn:     conn,
		exchange: exchange,
		logger:   logger,
	}, nil
}

func (s *Subscriber) Subscribe(c context.Context) (chan *carpb.CarEntity, func(), error) {
	masCh, cleanUp, err := s.SubscribeRaw(c)
	if err != nil {
		return nil, cleanUp, err
	}
	carCh := make(chan *carpb.CarEntity)
	go func() {
		for msg := range masCh {
			var car carpb.CarEntity
			err := json.Unmarshal(msg.Body, &car)
			if err != nil {
				s.logger.Error("cannot unmarshal", zap.Error(err))
			}
			carCh <- &car
		}
		close(carCh)
	}()
	return carCh, cleanUp, nil

}

func (s *Subscriber) SubscribeRaw(c context.Context) (<-chan amqp.Delivery, func(), error) {
	ch, err := s.conn.Channel()
	if err != nil {
		return nil, func() {}, fmt.Errorf("cannot allocate channle,%v", err)
	}
	closeCh := func() {
		err := ch.Close()
		if err != nil {
			s.logger.Error("cannot close channel", zap.Error(err))
		}
	}
	q, err := ch.QueueDeclare(
		"",    //name
		false, //furable
		true,  // autoDelete
		false, //exclusive
		false, // nowait
		nil,   // args
	)

	if err != nil {
		return nil, closeCh, fmt.Errorf("cannot declare queue: %v", err)
	}

	cleanUp := func() {
		_, err := ch.QueueDelete(
			q.Name,
			false, //
			false,
			false,
		)
		if err != nil {
			s.logger.Error("cannot QueueDelete", zap.Error(err))
		}
		closeCh()
	}
	ch.QueueBind(
		q.Name,
		"", //key
		s.exchange,
		false, //nowait
		nil,   // args
	)

	if err != nil {
		return nil, cleanUp, fmt.Errorf("cannot bind queue: %v", err)
	}
	msgs, err := ch.Consume(
		q.Name,
		"",    // consumer
		true,  //autoAck
		false, // exclusive
		false, //noLocal
		false, //noWarit
		nil,   // args
	)
	if err != nil {
		return nil, cleanUp, fmt.Errorf("cannot consume queue: %v", err)
	}
	return msgs, cleanUp, nil
}

func declareExchange(ch *amqp.Channel, exchange string) error {
	return ch.ExchangeDeclare(
		exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
}
