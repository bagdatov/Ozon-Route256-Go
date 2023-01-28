package repository

import (
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/bagdatov/homework-3/service-monitoring-order/internal/entity"
)

type producer struct {
	kafka sarama.SyncProducer
	reset string
}

func NewSyncProducer(brokers []string, resetTopic string) (*producer, error) {
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true

	prod, err := sarama.NewSyncProducer(brokers, conf)
	if err != nil {
		return nil, err
	}

	return &producer{
		kafka: prod,
		reset: resetTopic,
	}, nil
}

func (p *producer) Cancel(order entity.CancelOrder) error {

	o, err := json.Marshal(order)
	if err != nil {
		return err
	}

	_, _, err = p.kafka.SendMessage(&sarama.ProducerMessage{
		Topic: p.reset,
		Key:   sarama.StringEncoder(fmt.Sprint(order.OrderID)),
		Value: sarama.ByteEncoder(o),
	})
	if err != nil {
		return err
	}

	return nil
}
