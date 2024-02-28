package main

import (
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/wvalencia19/coffee-shop/models"
	"github.com/wvalencia19/coffee-shop/models/dtos"
)

func main() {
	config, err := dtos.ReadConfig("./data/config.json")
	if err != nil {
		logrus.Fatalf("Error reading config: %v", err)
	}

	workSimulator := models.CoffeeShopWorkSimulator{}
	shop := models.NewCoffeeShop(models.JSONGrindersTOGrinders(config.Grinders), models.JSONBrewersTOGBrewers(config.Brewers), config.NumBaristas, workSimulator)

	var wg sync.WaitGroup

	wg.Add(len(config.Orders))

	orderDone := make(chan int, 100)
	orders := models.JSONOrdersToOrders(config.Orders)
	go shop.StartProcessingOrders(&wg, orderDone)
	go shop.PlaceOrders(orders, &wg)

	wg.Wait()

	close(orderDone)

	processedOrders := make(map[int]bool)
	for id := range orderDone {
		processedOrders[id] = true
	}

	if len(processedOrders) != len(config.Orders) {
		logrus.Error("mismatch in number of orders placed and processed")
	} else {
		logrus.Info("all orders processed successfully")
	}
}
