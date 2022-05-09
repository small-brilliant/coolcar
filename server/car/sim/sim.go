package sim

import (
	"context"
	aipb "coolcar/ai/api/gen/v1"
	carpb "coolcar/car/api/gen/v1"
	"coolcar/car/mq"
	"fmt"
	"time"

	"go.uber.org/zap"
)

type Controller struct {
	CarService    carpb.CarServiceClient
	Aiservice     aipb.AiServiceClient
	CarSubscriber mq.Subscriber
	PosSubscriber mq.Subscriber
	Logger        *zap.Logger
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
	// car
	carCh, carcleanUp, err := c.CarSubscriber.Subscribe(ctx)
	defer carcleanUp()
	if err != nil {
		c.Logger.Error("cannot subscribe car", zap.Error(err))
		return
	}
	// pos
	posCh, poscleanUp, err := c.PosSubscriber.Subscribe(ctx)
	defer poscleanUp()
	if err != nil {
		c.Logger.Error("cannot subscribe position", zap.Error(err))
		return
	}
	//2.gorutining和carid对应
	//car
	carChans := make(map[string]chan *carpb.Car)
	posChans := make(map[string]chan *carpb.Location)
	for _, car := range cars {
		carFanoutCh := make(chan *carpb.Car)
		carChans[car.Id] = carFanoutCh
		posFanoutCh := make(chan *carpb.Location)
		posChans[car.Id] = posFanoutCh
		go c.SimulateCar(context.Background(), car, carFanoutCh, posFanoutCh)
	}

	// 3. 收消息
	for {
		select {
		case carUpdate := <-carCh:
			ch := carChans[carUpdate.Id]
			if ch != nil {
				ch <- carUpdate.Car
			}
		case posUpdate := <-posCh:
			ch := posChans[posUpdate.Id]
			if ch != nil && posUpdate.Car.Position != nil {
				ch <- &carpb.Location{
					Latitude:  posUpdate.Car.Position.Latitude,
					Longitude: posUpdate.Car.Position.Longitude,
				}
			}
		}
	}
}

func (c *Controller) SimulateCar(ctx context.Context, initial *carpb.CarEntity, carCh chan *carpb.Car, posCh chan *carpb.Location) {
	car := initial
	c.Logger.Info("simulations car.", zap.String("id", car.Id))
	for {
		select {
		case update := <-carCh:
			if update.Status == carpb.CarStatus_UNLOCKING {
				updated, err := c.unlockCar(ctx, car)
				if err != nil {
					c.Logger.Error("cannot unlock car", zap.Error(err))
					break
				}
				car = updated
			} else if update.Status == carpb.CarStatus_LOCKING {
				updated, err := c.lockCar(ctx, car)
				if err != nil {
					c.Logger.Error("cannot lock car", zap.Error(err))
					break
				}
				car = updated
			}
		case pos := <-posCh:
			updated, err := c.moveCar(ctx, car, pos)
			if err != nil {
				c.Logger.Error("cannot move car", zap.Error(err))
				break
			}
			car = updated
		}
	}
}

func (c *Controller) lockCar(ctx context.Context, car *carpb.CarEntity) (*carpb.CarEntity, error) {
	car.Car.Status = carpb.CarStatus_LOCKED
	_, err := c.CarService.UpdateCar(ctx, &carpb.UpdateCarRequest{
		Id:     car.Id,
		Status: car.Car.Status,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot lock car:%v", err)
	}
	_, err = c.Aiservice.EndSimulateCarPos(ctx, &aipb.EndSimulateCarPosRequest{
		CarId: car.Id,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot EndSimulateCarPos:%v", err)
	}
	return car, nil
}

func (c *Controller) unlockCar(ctx context.Context, car *carpb.CarEntity) (*carpb.CarEntity, error) {
	car.Car.Status = carpb.CarStatus_UNLOCKED
	_, err := c.CarService.UpdateCar(ctx, &carpb.UpdateCarRequest{
		Id:     car.Id,
		Status: car.Car.Status,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot Unlock car:%v", err)
	}
	_, err = c.Aiservice.SimulateCarPos(ctx, &aipb.SimulateCarPosRequest{
		CarId: car.Id,
		Pos: &aipb.Location{
			Latitude:  car.Car.Position.Latitude,
			Longitude: car.Car.Position.Longitude,
		},
	})
	if err != nil {
		c.Logger.Error("cannot SimulateCarPos", zap.Error(err))
	}
	return car, nil
}
func (c *Controller) moveCar(ctx context.Context, car *carpb.CarEntity, pos *carpb.Location) (*carpb.CarEntity, error) {
	car.Car.Position = pos
	_, err := c.CarService.UpdateCar(ctx, &carpb.UpdateCarRequest{
		Id:       car.Id,
		Position: pos,
	})
	if err != nil {
		return nil, fmt.Errorf("cannot update car:%v", err)
	}
	return car, nil
}
