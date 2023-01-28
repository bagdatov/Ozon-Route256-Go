package app

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"gitlab.ozon.dev/bagdatov/homework-2/api/pb"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/config"
	grpcServer "gitlab.ozon.dev/bagdatov/homework-2/game-master/delivery/grpc"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/delivery/http"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/migration"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/repository/postgres"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/repository/resty"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/usecase"
	"google.golang.org/grpc"
	"log"
	"net"
)

// Run -
func Run(conf *config.Config) {

	// migrations
	m, err := migration.New(conf)
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

	log.Println("Creating repositories...")
	// repository
	cli, err := resty.New(conf)
	if err != nil {
		log.Fatalf("failed to create parser: %v", err)
	}

	db, err := postgres.New(pool)
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}
	log.Println("Created")

	// use case
	log.Println("Creating usecase...")

	uc, err := usecase.New(cli, db)
	if err != nil {
		log.Fatalf("failed to create usecase: %v", err)
	}
	log.Println("Created")

	// delivery
	// grpc server
	log.Println("Creating grpc server...")

	newServer, err := grpcServer.New(uc)
	if err != nil {
		log.Fatalf("failed to create grpc delivery server: %v", err)
	}

	lis, err := net.Listen("tcp", conf.GRPC.Host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	server := grpc.NewServer(opts...)
	pb.RegisterChgkServer(server, newServer)
	log.Println("Server created")

	log.Println("gateway server starting...")
	// http gateway
	gatewayServer, err := http.NewGateway(conf.GRPC.Host, conf.Gateway.Port)
	if err != nil {
		log.Fatalf("failed to create gateway server: %v", err)
	}

	go func() {
		if err := gatewayServer.Server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	log.Println("gateway server started")

	log.Println("Ready")
	// run grpc server
	if err := server.Serve(lis); err != nil {
		log.Fatalf("gRPC server error: %s", err)
	}

	//TODO:: add graceful shutdown

}
