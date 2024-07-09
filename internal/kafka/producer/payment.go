package producer

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"payment-payments-api/internal/config"
	"payment-payments-api/internal/kafka/dto"
)

type PaymentProducer interface {
	Produce(message dto.PaymentRequest) error
}

func NewPaymentProducer(cfg *config.Config) (*paymentProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": cfg.BootstrapServers})
	if err != nil {
		return nil, err
	}
	return &paymentProducer{producer: p, topic: cfg.ProducerTopic}, nil
}

type paymentProducer struct {
	producer *kafka.Producer
	topic    string
}

func (p *paymentProducer) Produce(message dto.PaymentRequest) error {
	value, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %w", err)
	}

	deliveryChan := make(chan kafka.Event)
	defer close(deliveryChan)

	err = p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &p.topic, Partition: kafka.PartitionAny},
		Value:          value,
	}, deliveryChan)
	if err != nil {
		return fmt.Errorf("failed to produce message: %w", err)
	}

	e := <-deliveryChan
	msg := e.(*kafka.Message)
	if msg.TopicPartition.Error != nil {
		return fmt.Errorf("delivery failed: %v", msg.TopicPartition.Error)
	}

	log.Printf("Delivered message to %v", msg.TopicPartition)
	return nil
}
