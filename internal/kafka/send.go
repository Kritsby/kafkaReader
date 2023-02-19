package kafka

import (
	"dev/hak/internal/postgres"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
	"strconv"
	"strings"
	"time"
)

type SenderKafka struct {
	Repo postgres.Sender

	Consumer *kafka.Consumer
}

func NewSenderKafka(repo postgres.Sender, c *kafka.Consumer) *SenderKafka {
	return &SenderKafka{
		Repo: repo,

		Consumer: c,
	}
}

func (s *SenderKafka) TakeAndSend() error {
	for {
		ev := s.Consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			fmt.Printf("Message on %s\n", string(e.Value))
			str := fmt.Sprintf("%s", e.Value)

			dateFromKafka := strings.Join([]string{
				str[12:16], str[17:19], str[20:22], str[23:25], str[26:28]}, "")

			curDate := time.Now().UTC().Format("200601021504")

			dateFromKafkaInt, err := strconv.Atoi(dateFromKafka)
			if err != nil {
				return err
			}

			curDateInt, err := strconv.Atoi(curDate)
			if err != nil {
				return err
			}

			idStr := curDate
			if dateFromKafkaInt < curDateInt {
				idStr = dateFromKafka
			}

			idInt, err := strconv.Atoi(idStr)
			if err != nil {
				return err
			}

			s.parse(str[41:])

			err = s.Repo.Set(idInt, str)
			if err != nil {
				return err
			}
		case *kafka.Error:
			log.Println("kafka error", e.Error())
			return fmt.Errorf(e.Error())
		}

	}

	return nil
}

func (s *SenderKafka) parse(str string) {
	sliceStr := strings.Split(str, ",")
	fmt.Println(sliceStr)
}
