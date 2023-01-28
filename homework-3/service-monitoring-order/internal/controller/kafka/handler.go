package kafka

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"gitlab.ozon.dev/bagdatov/homework-3/service-monitoring-order/config"
	"gitlab.ozon.dev/bagdatov/homework-3/service-monitoring-order/internal/entity"
	"gitlab.ozon.dev/bagdatov/homework-3/service-monitoring-order/internal/usecase"
	"log"
	"time"
)

type kafkaHandler struct {
	uc   usecase.Monitoring
	conf config.Kafka
}

func NewHandler(uc usecase.Monitoring, conf config.Kafka) (*kafkaHandler, error) {
	return &kafkaHandler{
		uc:   uc,
		conf: conf,
	}, nil
}

func (c *kafkaHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *kafkaHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *kafkaHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {

	for msg := range claim.Messages() {

		log.Printf("consumer topic: <%v> msg: %s", msg.Topic, msg.Value)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

		switch msg.Topic {
		case c.conf.IncomeTopic:

			var o entity.Order

			err := json.Unmarshal(msg.Value, &o)
			if err != nil {
				log.Printf("incomingHandler.Unmarshal income error: %v", err)
				cancel()
				continue
			}

			err = c.uc.AddOrder(ctx, o)
			if err != nil {
				log.Printf("incomingHandler.AddOrder error: %v", err)
			}

		case c.conf.ResetTopic:

			var o entity.CancelOrder

			err := json.Unmarshal(msg.Value, &o)
			if err != nil {
				log.Printf("kafkaHandler.Unmarshal reset error: %v", err)
				cancel()
				continue
			}

			err = c.uc.Cancel(ctx, o.OrderID)
			if err != nil {
				log.Printf("resetHandler.Cancel error: %v", err)
			}

		case c.conf.ReservationTopic:

			var r entity.Reservation

			err := json.Unmarshal(msg.Value, &r)
			if err != nil {
				log.Printf("kafkaHandler.Unmarshal reservation error: %v", err)
				cancel()
				continue
			}

			err = c.uc.MarkReservation(ctx, r)
			if err != nil {
				log.Printf("kafkaHandler.MarkReservation error: %v", err)
			}
		}

		cancel()
	}

	return nil
}
