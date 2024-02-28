package models

import (
	"sync"

	"github.com/sirupsen/logrus"
)

type Barista struct {
	id int
}

func (b *Barista) ProcessOrders(shop *CoffeeShop, wg *sync.WaitGroup, orderDone chan<- int) {
	for order := range shop.orderQueue {
		coffee := shop.MakeCoffee(order)
		logrus.Infof("Barista %d completed order ID %d: Coffee of %d ounces", b.id, order.id, coffee.sizeOunces)
		orderDone <- order.id
		wg.Done()
	}
}

func NewBaristaList(numBaristas int) []*Barista {
	baristas := make([]*Barista, numBaristas)
	for i := range baristas {
		baristas[i] = &Barista{id: i}
	}

	return baristas
}
