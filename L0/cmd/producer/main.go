package main

import (
	k "demo_service/internal/kafka"
	"demo_service/internal/models"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
)

const (
	topic = "some_topic"
)

var paths = []string{"testdata/order1.json", "testdata/order2.json", "testdata/order3.json"}

func main() {
	log := setupLogger()

	producer, err := k.NewProducer("kafka:29091", log)
	if err != nil {
		log.Error("producer creation failed")
		os.Exit(1)
	}
	defer producer.Close()

	orders := readOrders()

	// time.Sleep(3)

	for _, order := range orders {
		err = producer.Produce(order, topic)

		if err != nil {
			fmt.Println("Order producing error")
		}
	}
	return
}

func readOrders() []models.Order {
	var orders []models.Order

	for _, path := range paths {
		data, err := os.ReadFile(path)

		if err != nil {
			fmt.Println("reading json from file error")
			os.Exit(1)
		}

		var order models.Order
		err = json.Unmarshal(data, &order)

		if err != nil {
			fmt.Println("json unmarshalling error")
		}

		fmt.Printf("Successfully parsed order: %s\n", order.OrderUId)

		orders = append(orders, order)
	}

	return orders
}

func setupLogger() (log *slog.Logger) {
	log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	return
}
