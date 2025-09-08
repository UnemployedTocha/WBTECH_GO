package kafka

import (
	"fmt"
	"log/slog"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Handler interface {
	HandleMessage(msg []byte) error
}

type Consumer struct {
	consumer *kafka.Consumer
	handler  Handler
	isStoped bool
	log      *slog.Logger
}

func NewConsumer(address string, handler Handler, topic, groupId string, logger *slog.Logger) (*Consumer, error) {
	cfg := kafka.ConfigMap{
		"bootstrap.servers":        address,
		"session.timeout.ms":       10000,
		"group.id":                 groupId,
		"enable.auto.commit":       true,
		"enable.auto.offset.store": false,
		"auto.offset.reset":        "earliest",
	}

	for {
		consumer, err := kafka.NewConsumer(&cfg)

		if err != nil {
			logger.Error("consumer creating error: %w | waiting 3 sec", slog.Any("err", err))
			time.Sleep(3 * time.Second)
			continue
		}

		if _, err = consumer.GetMetadata(&address, false, 3000); err != nil {
			logger.Error("consumer pinging error: %w | waiting 3 sec", slog.Any("err", err))
			time.Sleep(3 * time.Second)
			continue
		}

		if err = consumer.Subscribe(topic, nil); err != nil {
			logger.Error("consumer subscribe error: %w | waiting 3 sec", slog.Any("err", err))
			time.Sleep(3 * time.Second)
			continue
		}

		return &Consumer{consumer: consumer, handler: handler, isStoped: false, log: logger}, nil
	}
}

func (c *Consumer) Start() {

	c.log.Info("starting consumer")

	for {
		if c.isStoped == true {
			return
		}

		msg, err := c.consumer.ReadMessage(-1)

		if err != nil {
			c.log.Error("reading kafka message error", slog.Any("err", err))
		}

		if msg == nil {
			c.log.Error("Received nil message, continue consumer work")
			continue
		}

		if err = c.handler.HandleMessage(msg.Value); err != nil {
			c.log.Error("handling message error: %w")
			continue
		}

		c.log.Info("message handled!!", slog.Any("val: ", msg.Value))

		if _, err = c.consumer.StoreMessage(msg); err != nil {
			c.log.Error("kafka store offset failed")
		}

	}
}

func (c *Consumer) Stop() error {
	c.isStoped = true
	if _, err := c.consumer.Commit(); err != nil {
		return fmt.Errorf("error commiting while stopping error: %w", err)
	}
	return c.consumer.Close()
}
