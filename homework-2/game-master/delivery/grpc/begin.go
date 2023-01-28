package grpc

import (
	"context"
	"log"

	"gitlab.ozon.dev/bagdatov/homework-2/api/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s myServer) Begin(ctx context.Context, message *pb.GameRequest) (*pb.GameResponse, error) {
	if message.GetTournament() == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid tournament key")
	}

	if message.GetChatID() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid chat id")
	}

	qs, err := s.game.Begin(ctx, message.GetTournament(), message.GetChatID())
	if err != nil {
		log.Println(err, message)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GameResponse{
		QuestionIDs: qs,
	}, nil
}
