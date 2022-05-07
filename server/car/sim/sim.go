package sim

import (
	"context"
	carpb "coolcar/car/api/gen/v1"
	"coolcar/car/mq"
	"time"

	"go.uber.org/zap"
)

type Controller struct {
	CarService carpb.CarServiceClient
	Subscriber mq.Subscriber
	Logger     *zap.Logger
}

func (c *Controller) RunSimulations(ctx context.Context) {
	//针对每一台汽车模拟
	//1.获取所有车
	var cars []*carpb.CarEntity
	for {
		time.Sleep(3 * time.Second)
		res, err := c.CarService.GetCars(ctx, &carpb.GetCarsRequest{})
		if err != nil {
			c.Logger.Error("cannot get cars", zap.Error(err))
			continue
		}
		cars = res.Cars
		break
	}

	c.Logger.Info("Runing car simulations.", zap.Int("car_count", len(cars)))
	msgCh, cleanUp, err := c.Subscriber.Subscribe(ctx)
	defer cleanUp()
	if err != nil {
		c.Logger.Error("cannot subscribe", zap.Error(err))
		return
	}
	//2.gorutining和carid对应
	carChans := make(map[string]chan *carpb.Car)
	for _, car := range cars {
		ch := make(chan *carpb.Car)
		carChans[car.Id] = ch
		go c.SimulateCar(context.Background(), car, ch)
	}

	// 3. 收消息
	for carUpdate := range msgCh {
		ch := carChans[carUpdate.Id]
		if ch != nil {
			ch <- carUpdate.Car
		}
	}
}

func (c *Controller) SimulateCar(ctx context.Context, initial *carpb.CarEntity, ch chan *carpb.Car) {
	carID := initial.Id
	c.Logger.Info("simulations car.", zap.String("id", carID))

	for update := range ch {
		if update.Status == carpb.CarStatus_UNLOCKING {
			_, err := c.CarService.UpdateCar(ctx, &carpb.UpdateCarRequest{
				Id:     carID,
				Status: carpb.CarStatus_UNLOCKED,
			})
			if err != nil {
				c.Logger.Error("cannot unlock car", zap.Error(err))
			}
		} else if update.Status == carpb.CarStatus_LOCKING {
			_, err := c.CarService.UpdateCar(ctx, &carpb.UpdateCarRequest{
				Id:     carID,
				Status: carpb.CarStatus_LOCKED,
			})
			if err != nil {
				c.Logger.Error("cannot lock car", zap.Error(err))
			}
		}
	}
}
