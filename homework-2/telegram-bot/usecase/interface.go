package usecase

import (
	"context"
	"gitlab.ozon.dev/bagdatov/homework-2/telegram-bot/models"
)

type Bot interface {
	Begin(ctx context.Context, chatID int64, tournament string) string
	Submit(ctx context.Context, chatID int64, replyTo int, username, guess string) string
	Details(ctx context.Context, tournamentKey string) string
	Score(ctx context.Context, chatID int64) string
	Random(ctx context.Context) string
	Next(chatID int64) string
	Stop(chatID int64) string
	Help() string

	Question(chatID, questionID int64) int
	Answer(chatID, questionID int64)
	FinishTournament(chatID int64)
}

type Game interface {
	Tournament(ctx context.Context, tournamentKey string) (models.Tournament, error)
	Begin(ctx context.Context, chatID int64, tournament string) ([]int64, error)
	Submit(ctx context.Context, chatID, questionID int64, username, answer string) (bool, error)
	Question(ctx context.Context, questionID int64) (models.Question, error)
	Answer(ctx context.Context, questionID int64) (models.Question, error)
	Score(ctx context.Context, chatID int64) ([]models.User, error)
	Random(ctx context.Context) ([]models.Tournament, error)
	Finish(ctx context.Context, chatID int64) error
	//Sessions(ctx context.Context, chatID int64) ([]models.Session, error)
}

type Messenger interface {
	Send(chatID int64, text string) (int, error)
}
