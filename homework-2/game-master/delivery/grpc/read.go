package grpc

import (
	"context"
	"log"

	"gitlab.ozon.dev/bagdatov/homework-2/api/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s myServer) ReadTournament(ctx context.Context, message *pb.TournamentRequest) (*pb.Tournament, error) {
	if message.GetName() == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid tournament key")
	}

	t, err := s.game.FindTournament(ctx, message.GetName())
	if err != nil {
		log.Println(err, message)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Tournament{
		ID:          t.ID,
		Key:         t.TextId,
		Name:        t.Title,
		Date:        t.CreatedAt,
		ToursNum:    t.ChildrenNum,
		QuestionNum: t.QuestionsNum,
	}, nil
}

func (s myServer) ReadQuestion(ctx context.Context, message *pb.QuestionRequest) (*pb.Question, error) {
	if message.GetQuestionID() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid question id")
	}

	q, err := s.game.FetchQuestion(ctx, message.GetQuestionID())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Question{
		Num:  q.Number,
		Text: q.Question,
	}, nil
}

func (s myServer) ReadAnswer(ctx context.Context, message *pb.AnswerRequest) (*pb.Answer, error) {
	if message.GetQuestionID() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid question id")
	}

	q, err := s.game.FetchQuestion(ctx, message.GetQuestionID())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Answer{
		Num:     q.Number,
		Text:    q.Answer,
		Comment: q.Comments,
		Source:  q.Sources,
		Author:  q.Authors,
	}, nil
}

func (s myServer) ReadScore(ctx context.Context, message *pb.ScoreRequest) (*pb.Score, error) {
	if message.GetChatID() == 0 {
		return nil, status.Error(codes.InvalidArgument, "invalid chat id")
	}

	scores, err := s.game.FetchScore(ctx, message.GetChatID())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	users := make([]*pb.User, 0, len(scores))

	for _, s := range scores {
		u := &pb.User{
			Name:  s.Name,
			Score: s.Score,
		}
		users = append(users, u)
	}

	return &pb.Score{
		Users: users,
	}, nil
}
