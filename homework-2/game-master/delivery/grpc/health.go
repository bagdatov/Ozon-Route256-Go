package grpc

import (
	"context"
	"fmt"
	"gitlab.ozon.dev/bagdatov/homework-2/api/pb"
)

func (s myServer) HealthCheck(ctx context.Context, message *pb.Ping) (*pb.Pong, error) {
	return &pb.Pong{
		ChatID: message.GetChatID(),
		Data:   fmt.Sprintf("@%s, Hello and pong", message.GetUsername()),
	}, nil
}
