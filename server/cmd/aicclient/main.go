package main

import (
	"context"
	aipb "coolcar/ai/api/gen/v1"
	carpb "coolcar/car/api/gen/v1"
	"coolcar/car/mq/amqpclient"
	"encoding/json"
	"fmt"
	"time"

	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:18001", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	ac := aipb.NewAiServiceClient(conn)
	c := context.Background()

	// Measure distance.
	res, err := ac.MeasureDistance(c, &aipb.MeasureDistanceRequest{
		From: &aipb.Location{
			Latitude:  29.756825521115363,
			Longitude: 121.87222114786053,
		},
		To: &aipb.Location{
			Latitude:  29.757211315878838,
			Longitude: 122.87024571958649,
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("行驶距离：%+v\n", res.DistanceKm)

	// Licsense recognition.
	idRes, err := ac.LicIdentity(c, &aipb.LicIdentityRequest{
		Photo: []byte{1, 2, 3, 4, 5},
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", idRes)

	// Car position simulation.
	_, err = ac.SimulateCarPos(c, &aipb.SimulateCarPosRequest{
		CarId: "car123",
		Pos: &aipb.Location{
			Latitude:  30,
			Longitude: 120,
		},
	})
	if err != nil {
		panic(err)
	}

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	amqpConn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}

	sub, err := amqpclient.NewSubscriber(amqpConn, "pos_sim", logger)
	if err != nil {
		panic(err)
	}

	ch, cleanUp, err := sub.SubscribeRaw(c)
	defer cleanUp()

	if err != nil {
		panic(err)
	}

	tm := time.After(10 * time.Second)
	i := 0
	for {
		shouldStop := false
		select {
		case msg := <-ch:
			var update carpb.CarEntity
			err = json.Unmarshal(msg.Body, &update)
			if err != nil {
				panic(err)
			}
			i += 1
			fmt.Printf("第%d更新：%+v\n", i, &update)

		case <-tm:
			shouldStop = true
		}
		if shouldStop {
			break
		}
	}

	_, err = ac.EndSimulateCarPos(c, &aipb.EndSimulateCarPosRequest{
		CarId: "car123",
	})
	if err != nil {
		panic(err)
	}
}
