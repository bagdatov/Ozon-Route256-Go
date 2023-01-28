package usecase

import (
	"context"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/models"
)

type Game interface {
	Begin(ctx context.Context, tournamentKey string, chatID int64) (questions []int64, err error)
	FindTournament(ctx context.Context, tournamentKey string) (models.Tournament, error)
	FetchQuestion(ctx context.Context, questionID int64) (models.Question, error)
	FetchScore(ctx context.Context, chatID int64) ([]models.User, error)
	FetchSessions(ctx context.Context, chatID int64) ([]models.Session, error)
	FinishSession(ctx context.Context, chatID int64) error
	RandomTournaments(ctx context.Context) ([]models.TourI, error)
	SubmitGuess(ctx context.Context, chatID, questionID int64, username, guess string) (bool, error)
}

type Parser interface {
	Tournament(ctx context.Context, name string) (models.Tournament, error)
	Tour(ctx context.Context, name string) (models.Tour, error)
	Root(ctx context.Context) (models.Tournament, error)
}

type Database interface {
	AddTournament(ctx context.Context, tournament models.Tournament, tours []models.Tour) error
	AddScore(ctx context.Context, chatID, questionID int64, user string, isCorrect bool) error
	AddSession(ctx context.Context, tournamentID int64, chatID int64) error
	FetchQuestion(ctx context.Context, questionID int64) (models.Question, error)
	FetchQuestions(ctx context.Context, tournamentID int64) (questions []int64, err error)
	FetchScore(ctx context.Context, chatID int64) ([]models.User, error)
	FetchSessions(ctx context.Context, chatID int64) ([]models.Session, error)
	FinishSession(ctx context.Context, chatID int64) error
	IsCached(ctx context.Context, tournamentName string) (cashed bool, tournamentID int64, err error)
}
