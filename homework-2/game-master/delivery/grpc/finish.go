package grpc

import (
	"context"
	"log"

	"gitlab.ozon.dev/bagdatov/homework-2/api/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s myServer) Finish(ctx context.Context, message *pb.StopRequest) (*pb.StopResponse, error) {
	if message.GetChatID() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid chat id")
	}

	err := s.game.FinishSession(ctx, message.GetChatID())
	if err != nil {
		log.Println(err, message)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.StopResponse{}, nil
}
