package kafka

import (
	"demo_service/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var UnknownEventError = errors.New("unknown event type")

type Producer struct {
	producer *kafka.Producer
	log      *slog.Logger
}

func NewProducer(address string, logger *slog.Logger) (*Producer, error) {
	cfg := kafka.ConfigMap{
		"bootstrap.servers": address,
	}

	logger.Info("making consumer")

	for {
		p, err := kafka.NewProducer(&cfg)
		if err != nil {
			logger.Error("kafka producer creation failed, waiting 3 sec, ", slog.Any("err: ", err))
			time.Sleep(3 * time.Second)
			continue
		}

		_, err = p.GetMetadata(nil, true, 3000)
		if err != nil {
			p.Close()
			logger.Error("kafka producer pinging failed, waiting 3 sec, ", slog.Any("err: ", err))
			continue
		}
		return &Producer{producer: p, log: logger}, nil
	}
}

func (p *Producer) Produce(order models.Order, topic string) error {
	msg, err := json.Marshal(order)
	if err != nil {
		return fmt.Errorf("order json decoding error: %w", err)
	}

	kafkaChan := make(chan kafka.Event)
	err = p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msg,
	}, kafkaChan)

	if err != nil {
		return fmt.Errorf("kafka produce error: %w", err)
	}

	event := <-kafkaChan
	switch er := event.(type) {
	case *kafka.Message:
		{
			fmt.Println("msg successfully sent!!!1!")
			return nil
		}
	case *kafka.Error:
		return fmt.Errorf("kafka response error: %w", er)
	default:
		return UnknownEventError
	}
}

func (p *Producer) Close() {
	timeout := 3000
	p.producer.Flush(timeout)
	p.producer.Close()
}
