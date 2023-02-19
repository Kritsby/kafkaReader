package app

import (
	"dev/hak/internal/config"
	"dev/hak/internal/kafka"
	"dev/hak/internal/postgres"
	"log"
)

func Run(cfg config.Config) error {
	pool, err := postgres.NewPostgresql(cfg.PSQL)
	if err != nil {
		return err
	}

	c, err := kafka.NewKafkaConn(cfg.Kafka)
	if err != nil {
		return err
	}
	defer c.Consumer.Close()

	repo := postgres.NewRepository(pool)
	k := kafka.NewKafkaConsumer(repo, c.Consumer)

	if err = k.Consumerer.TakeAndSend(); err != nil {
		log.Println(err)
	}

	return nil
}
