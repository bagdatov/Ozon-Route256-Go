package repository

import (
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/bagdatov/homework-3/service-store/internal/entity"
)

type producer struct {
	kafka       sarama.SyncProducer
	reset       string
	reservation string
}

func NewSyncProducer(brokers []string, resetTopic, reservationTopic string) (*producer, error) {
	conf := sarama.NewConfig()
	conf.Producer.Return.Successes = true

	prod, err := sarama.NewSyncProducer(brokers, conf)
	if err != nil {
		return nil, err
	}

	return &producer{
		kafka:       prod,
		reset:       resetTopic,
		reservation: reservationTopic,
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

func (p *producer) Publish(reservation entity.Reservation) error {

	o, err := json.Marshal(reservation)
	if err != nil {
		return err
	}

	_, _, err = p.kafka.SendMessage(&sarama.ProducerMessage{
		Topic: p.reservation,
		Key:   sarama.StringEncoder(fmt.Sprint(reservation.OrderID)),
		Value: sarama.ByteEncoder(o),
	})
	if err != nil {
		return err
	}

	return nil
}
