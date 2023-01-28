package app

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/bagdatov/homework-3/service-create-order/config"
	"gitlab.ozon.dev/bagdatov/homework-3/service-create-order/internal/controller/http"
	"gitlab.ozon.dev/bagdatov/homework-3/service-create-order/internal/repository"
	"gitlab.ozon.dev/bagdatov/homework-3/service-create-order/internal/usecase"
	"gitlab.ozon.dev/bagdatov/homework-3/service-create-order/migrations"
	"log"
)

func Run(conf *config.Config) {

	// migrations
	m, err := migrations.New(conf)
	if err != nil {
		log.Fatalf("failed to create migration: %v", err)
	}

	if err := m.RunBlocking(); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("Connecting to pg...")

	pool, err := pgxpool.Connect(ctx, fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=public",
		conf.PG.Username, conf.PG.Password, conf.PG.Host, conf.PG.Port, conf.PG.Dbname))
	if err != nil {
		log.Fatalf("failed to connect to pg: %v", err)
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("failed to ping postgres: %v", err)
	}

	log.Println("Connected")

	// repositories
	log.Println("Creating repositories...")
	db := repository.NewDB(pool)

	broker, err := repository.NewSyncProducer(conf.Brokers, conf.IncomeTopic, conf.ResetTopic)
	if err != nil {
		log.Fatalf("Failed to connect to kafka: %v", err)
	}

	log.Println("Created")

	// use case
	uc := usecase.New(db, broker)

	// controller
	c := http.NewController(uc)

	// start server
	log.Fatal(
		c.Start(conf.App.Host),
	)
}
