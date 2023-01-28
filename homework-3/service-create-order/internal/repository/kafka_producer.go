package repository

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/bagdatov/homework-3/service-create-order/internal/entity"
)

type producer struct {
	kafka  sarama.SyncProducer
	income string
	reset  string
}

func NewSyncProducer(brokers []string, incomeTopic, resetTopic string) (*producer, error) {
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true

	prod, err := sarama.NewSyncProducer(brokers, conf)
	if err != nil {
		return nil, err
	}

	return &producer{
		kafka:  prod,
		income: incomeTopic,
		reset:  resetTopic,
	}, nil
}

func (p *producer) Publish(order entity.Order) error {

	o, err := json.Marshal(order)
	if err != nil {
		return err
	}

	_, _, err = p.kafka.SendMessage(&sarama.ProducerMessage{
		Topic: p.income,
		Key:   sarama.StringEncoder(fmt.Sprint(order.ID)),
		Value: sarama.ByteEncoder(o),
	})
	if err != nil {
		return err
	}

	return nil
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
