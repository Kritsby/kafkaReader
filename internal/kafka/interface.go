package kafka

import (
	"dev/hak/internal/postgres"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Consumerer interface {
	TakeAndSend() error
}

type ConsumerKafka struct {
	Consumerer
}

func NewKafkaConsumer(repo *postgres.Repository, c *kafka.Consumer) *ConsumerKafka {
	return &ConsumerKafka{
		Consumerer: NewSenderKafka(repo, c),
	}
}
