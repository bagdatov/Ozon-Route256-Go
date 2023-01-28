package grpc

import (
	"context"
	"gitlab.ozon.dev/bagdatov/homework-2/api/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s myServer) Submit(ctx context.Context, m *pb.Guess) (*pb.GuessResponse, error) {
	if m.GetChatID() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid chat id")
	}

	if m.GetQuestionID() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid question id")
	}

	if m.GetAnswer() == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid answer - empty")
	}

	if m.GetUsername() == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid username")
	}

	isCorrect, err := s.game.SubmitGuess(ctx, m.GetChatID(), m.GetQuestionID(), m.GetUsername(), m.GetAnswer())
	if err != nil {
		log.Println(err, m)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.GuessResponse{
		Correct: isCorrect,
	}, nil
}
