package app

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"

	"github.com/Shopify/sarama"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/bagdatov/homework-3/service-store/config"
	"gitlab.ozon.dev/bagdatov/homework-3/service-store/internal/controller/http"
	"gitlab.ozon.dev/bagdatov/homework-3/service-store/internal/controller/kafka"
	"gitlab.ozon.dev/bagdatov/homework-3/service-store/internal/repository"
	"gitlab.ozon.dev/bagdatov/homework-3/service-store/internal/usecase"
	"gitlab.ozon.dev/bagdatov/homework-3/service-store/migrations"
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

	log.Println("Connecting to redis...")

	cli := redis.NewClient(&redis.Options{
		Addr: conf.Redis.Address,
		DB:   conf.Redis.DB,
	})

	if err := cli.Ping(ctx).Err(); err != nil {
		log.Fatalf("failed to ping redis: %v", err)
	}

	log.Println("Connected")

	// repositories
	log.Println("Creating repositories...")
	db := repository.NewDB(pool)

	broker, err := repository.NewSyncProducer(conf.Brokers, conf.ResetTopic, conf.ReservationTopic)
	if err != nil {
		log.Fatalf("Failed to connect to kafka producer: %v", err)
	}

	rd := repository.NewCache(cli, conf.Redis.TTL)

	log.Println("Created")

	// use case
	uc := usecase.New(db, db, broker, rd)

	// http controller
	c := http.NewController(uc)

	// kafka controller
	kConfig := sarama.NewConfig()
	kafka.Start(ctx, conf, kConfig, uc)

	// start server
	log.Fatal(
		c.Start(conf.App.Host),
	)
}
