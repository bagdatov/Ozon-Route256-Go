package http

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"gitlab.ozon.dev/bagdatov/homework-2/api/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

type Gateway struct {
	mux    *runtime.ServeMux
	conn   *grpc.ClientConn
	Server *http.Server
}

func NewGateway(target, port string) (*Gateway, error) {
	mux := runtime.NewServeMux()

	conn, err := grpc.DialContext(
		context.Background(),
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to dial http gateway: %w", err)
	}
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	err = pb.RegisterChgkHandler(context.Background(), mux, conn)
	if err != nil {
		return nil, fmt.Errorf("failed to register http handler: %w", err)
	}

	return &Gateway{
		mux:    mux,
		conn:   conn,
		Server: server,
	}, nil
}
