package kafka

import (
	"dev/hak/internal/config"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaConsumer struct {
	Consumer *kafka.Consumer
}

func NewKafkaConn(cfg config.Kafka) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  "rc1a-b5e65f36lm3an1d5.mdb.yandexcloud.net:9091",
		"sasl.mechanisms":    "SCRAM-SHA-512",
		"security.protocol":  "SASL_SSL",
		"sasl.username":      "9433_reader",
		"sasl.password":      "eUIpgWu0PWTJaTrjhjQD3.hoyhntiK",
		"ssl.ca.location":    "./ca.pem",
		"session.timeout.ms": 6000,
		"group.id":           "sabmission",
		"enable.auto.commit": false,
		"auto.offset.reset":  "smallest"})
	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{"zsmk-9433-dev-01"}, nil)

	return &KafkaConsumer{Consumer: c}, nil
}
