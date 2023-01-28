package kafka

import (
	"context"
	"log"
	"time"

	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/bagdatov/homework-3/service-store/config"
	"gitlab.ozon.dev/bagdatov/homework-3/service-store/internal/usecase"
)

func Start(ctx context.Context, conf *config.Config, kafkaConf *sarama.Config, uc usecase.Storage) {
	reserveHandler, err := NewHandler(uc, conf.Kafka)
	if err != nil {
		log.Fatalf("Failed to create kafka handler: %v", err)
	}

	group, err := sarama.NewConsumerGroup(conf.Kafka.Brokers, conf.Kafka.GroupID, kafkaConf)
	if err != nil {
		log.Fatalf("Failed to create kafka consumer group: %v", err)
	}

	go func() {
		for {
			err := group.Consume(ctx, []string{
				conf.Kafka.ReservationTopic,
				conf.Kafka.IncomeTopic,
				conf.Kafka.ResetTopic,
			}, reserveHandler)

			if err != nil {
				log.Printf("kafka consumer error: %v", err)
				time.Sleep(time.Second * 5)
			}
		}
	}()
}
