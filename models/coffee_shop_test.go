package models

import (
	"sync"
	"testing"
)

func TestCoffeeShop(t *testing.T) {
	numBaristas := 2
	grinders := []*Grinder{NewGrinder(10)}
	brewers := []*Brewer{NewBrewer(5)}
	mockSleeper := MockWorkSimulator{}

	shop := NewCoffeeShop(grinders, brewers, numBaristas, mockSleeper)

	var wg sync.WaitGroup
	orderDone := make(chan int, 100)

	orders := []Order{{ouncesOfCoffeeWanted: 12}, {ouncesOfCoffeeWanted: 8}}
	wg.Add(len(orders))

	shop.PlaceOrders(orders, &wg)

	shop.StartProcessingOrders(&wg, orderDone)

	wg.Wait()
	close(orderDone)

	processedCount := 0
	for range orderDone {
		processedCount++
	}

	if processedCount != len(orders) {
		t.Errorf("Expected %d orders processed, got %d", len(orders), processedCount)
	}
}

func TestMakeCoffee(t *testing.T) {
	grinders := []*Grinder{{id: 1, gramsPerSecond: 10, busy: false}}
	brewers := []*Brewer{{id: 1, ouncesWaterPerSecond: 5, busy: false}}
	shop := NewCoffeeShop(grinders, brewers, 1, &MockWorkSimulator{})

	order := Order{ouncesOfCoffeeWanted: 10}

	coffee := shop.MakeCoffee(order)

	// Since 2 grams of beans are needed per ounce of coffee wanted,
	// and the brewing process involves specific ratios and multipliers
	expectedOunces := ((gramsNeededPerOunce * order.ouncesOfCoffeeWanted) / gramsToOuncesRatio) * ouncesMultiplier

	if coffee.sizeOunces != expectedOunces {
		t.Errorf("Expected %d ounces of coffee, got %d", expectedOunces, coffee.sizeOunces)
	}
}
