package models

import "github.com/wvalencia19/coffee-shop/models/dtos"

type Order struct {
	id                   int
	ouncesOfCoffeeWanted int
}

func NewOrder(ouncesOfCoffeeWanted int) Order {
	return Order{
		ouncesOfCoffeeWanted: ouncesOfCoffeeWanted,
	}
}

func JSONOrdersToOrders(jo []dtos.Order) []Order {
	orders := make([]Order, len(jo))
	for i, order := range jo {
		orders[i] = NewOrder(order.OuncesOfCoffeeWanted)
	}

	return orders
}
