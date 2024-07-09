package consumer

import (
	"encoding/json"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"payment-payments-api/internal/config"
	"payment-payments-api/internal/kafka/dto"
	"payment-payments-api/internal/services"
)

type Consumer interface {
	Consume(s *services.Services)
}

func NewPaymentConsumer(cfg *config.Config) (*paymentConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.BootstrapServers,
		"group.id":          cfg.GroupID,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		return nil, err
	}

	return &paymentConsumer{consumer: c, topic: cfg.ConsumerTopic}, nil
}

type paymentConsumer struct {
	consumer *kafka.Consumer
	topic    string
}

func (c *paymentConsumer) Consume(s *services.Services) {
	defer c.consumer.Close()

	err := c.consumer.Subscribe(c.topic, nil)
	if err != nil {
		log.Fatalf("Error subscribing to topic: %v", err)
	}

	for {
		msg, err := c.consumer.ReadMessage(-1)
		if err == nil {
			var message dto.PaymentResponse
			err := json.Unmarshal(msg.Value, &message)
			if err != nil {
				log.Printf("Error unmarshalling message: %v", err)
				continue
			}
			s.Payment.UpdatePayment(message)

		} else {
			log.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}
}
