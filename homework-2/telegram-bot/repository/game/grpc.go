package game

import (
	"context"
	"gitlab.ozon.dev/bagdatov/homework-2/api/pb"
	"gitlab.ozon.dev/bagdatov/homework-2/telegram-bot/models"
)

type game struct {
	cli pb.ChgkClient
}

func New(cli pb.ChgkClient) *game {
	return &game{
		cli: cli,
	}
}

func (g *game) Begin(ctx context.Context, chatID int64, tournament string) ([]int64, error) {
	resp, err := g.cli.Begin(ctx, &pb.GameRequest{
		Tournament: tournament,
		ChatID:     chatID,
	})
	if err != nil {
		return nil, err
	}

	return resp.QuestionIDs, nil
}

func (g *game) Submit(ctx context.Context, chatID, questionID int64, username, answer string) (bool, error) {
	resp, err := g.cli.Submit(ctx, &pb.Guess{
		ChatID:     chatID,
		Username:   username,
		QuestionID: questionID,
		Answer:     answer,
	})
	if err != nil {
		return false, err
	}

	return resp.Correct, nil
}

func (g *game) Question(ctx context.Context, questionID int64) (models.Question, error) {

	resp, err := g.cli.ReadQuestion(ctx, &pb.QuestionRequest{
		QuestionID: questionID,
	})
	if err != nil {
		return models.Question{}, err
	}

	return models.Question{
		QuestionId: questionID,
		Number:     resp.GetNum(),
		Question:   resp.GetText(),
	}, nil
}

func (g *game) Answer(ctx context.Context, questionID int64) (models.Question, error) {

	resp, err := g.cli.ReadAnswer(ctx, &pb.AnswerRequest{
		QuestionID: questionID,
	})
	if err != nil {
		return models.Question{}, err
	}

	return models.Question{
		QuestionId: questionID,
		Number:     resp.GetNum(),
		Answer:     resp.GetText(),
		Comments:   resp.GetComment(),
		Authors:    resp.GetAuthor(),
	}, nil
}

func (g *game) Score(ctx context.Context, chatID int64) ([]models.User, error) {

	resp, err := g.cli.ReadScore(ctx, &pb.ScoreRequest{
		ChatID: chatID,
	})
	if err != nil {
		return nil, err
	}

	users := make([]models.User, 0, len(resp.Users))
	for _, u := range resp.Users {

		users = append(users, models.User{
			Name:  u.GetName(),
			Score: u.GetScore(),
		})
	}

	return users, nil

}

func (g *game) Finish(ctx context.Context, chatID int64) error {

	_, err := g.cli.Finish(ctx, &pb.StopRequest{
		ChatID: chatID,
	})

	return err

}

func (g *game) Tournament(ctx context.Context, tournamentKey string) (models.Tournament, error) {

	resp, err := g.cli.ReadTournament(ctx, &pb.TournamentRequest{
		Name: tournamentKey,
	})
	if err != nil {
		return models.Tournament{}, err
	}

	return models.Tournament{
		ID:          resp.GetID(),
		Key:         resp.GetKey(),
		ToursNum:    resp.GetToursNum(),
		QuestionNum: resp.GetQuestionNum(),
		Name:        resp.GetName(),
		Date:        resp.GetDate(),
	}, nil
}

func (g *game) Random(ctx context.Context) ([]models.Tournament, error) {

	resp, err := g.cli.RandomTournaments(ctx, &pb.RandomRequest{})
	if err != nil {
		return nil, err
	}

	tournaments := make([]models.Tournament, 0, len(resp.GetTournaments()))

	for _, t := range resp.GetTournaments() {
		tournaments = append(tournaments, models.Tournament{
			ID:          t.GetID(),
			Key:         t.GetKey(),
			ToursNum:    t.GetToursNum(),
			QuestionNum: t.GetQuestionNum(),
			Name:        t.GetName(),
			Date:        t.GetDate(),
		})
	}

	return tournaments, nil
}
