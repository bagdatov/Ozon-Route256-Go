package grpc

import (
	"context"
	"gitlab.ozon.dev/bagdatov/homework-2/api/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s myServer) RandomTournaments(ctx context.Context, message *pb.RandomRequest) (*pb.RandomResponse, error) {
	tournaments, err := s.game.RandomTournaments(ctx)
	if err != nil {
		log.Println(err, message)
		return nil, status.Error(codes.Internal, err.Error())
	}

	r := &pb.RandomResponse{
		Tournaments: make([]*pb.Tournament, 0, len(tournaments)),
	}

	for _, t := range tournaments {
		r.Tournaments = append(r.Tournaments, &pb.Tournament{
			ID:          t.ID,
			Key:         t.TextId,
			Name:        t.Title,
			Date:        t.CreatedAt,
			ToursNum:    t.ChildrenNum,
			QuestionNum: t.QuestionsNum,
		})
	}

	return r, nil
}
