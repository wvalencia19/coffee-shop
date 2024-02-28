package models

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	gramsNeededPerOunce = 2
	orderQueueSize      = 100
)

// WorkSimulator interface for abstracting time management
type WorkSimulator interface {
	SimulateWork(duration time.Duration)
}

// CoffeeShopWorkSimulator uses actual SimulateWork
type CoffeeShopWorkSimulator struct{}

func (rts CoffeeShopWorkSimulator) SimulateWork(duration time.Duration) {
	time.Sleep(duration)
}

// Mock MockWorkSimulator to avoid actual sleeping in tests
type MockWorkSimulator struct{}

// Do nothing
func (mtc MockWorkSimulator) SimulateWork(duration time.Duration) {

}

type Beans struct {
	weightGrams int
}

type Coffee struct {
	sizeOunces int
}

type CoffeeShop struct {
	baristas      []*Barista
	orderQueue    chan Order
	grinderPool   chan GrinderInterface
	brewerPool    chan BrewerInterface
	WorkSimulator WorkSimulator
}

func NewCoffeeShop(grinders []*Grinder, brewers []*Brewer, numBaristas int, WorkSimulator WorkSimulator) *CoffeeShop {
	grinderPool := NewGrinderPool(grinders)

	brewerPool := NewBrewerPool(brewers)

	baristas := NewBaristaList(numBaristas)

	return &CoffeeShop{
		baristas:      baristas,
		orderQueue:    make(chan Order, orderQueueSize),
		grinderPool:   grinderPool,
		brewerPool:    brewerPool,
		WorkSimulator: WorkSimulator,
	}
}

func (shop *CoffeeShop) MakeCoffee(order Order) Coffee {
	ungroundBeans := Beans{weightGrams: gramsNeededPerOunce * order.ouncesOfCoffeeWanted}

	grinder := <-shop.grinderPool
	groundBeans := grinder.Grind(ungroundBeans, shop.WorkSimulator)
	shop.grinderPool <- grinder

	brewer := <-shop.brewerPool
	coffee := brewer.Brew(groundBeans, shop.WorkSimulator)
	shop.brewerPool <- brewer

	return coffee
}

func (shop *CoffeeShop) StartProcessingOrders(wg *sync.WaitGroup, orderDone chan<- int) {
	for _, barista := range shop.baristas {
		go barista.ProcessOrders(shop, wg, orderDone)
	}
}

func (shop *CoffeeShop) PlaceOrders(orders []Order, wg *sync.WaitGroup) {
	for i, order := range orders {
		wg.Add(1)
		go func(orderNum int, order Order) {
			defer wg.Done()

			order.id = orderNum
			shop.orderQueue <- order
			logrus.Infof("Order ID %d placed\n", orderNum)
		}(i, order)
	}
}
