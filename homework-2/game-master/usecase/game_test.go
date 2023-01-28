package usecase

import (
	"context"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/models"
	"reflect"
	"testing"
	"time"
)

func TestGame_Begin(t *testing.T) {
	mc := minimock.NewController(t)
	defer mc.Finish()

	var (
		chatID        int64 = 1
		tournamentID  int64 = 72491
		tournamentKey       = "olivje22_u"
		questions           = []int64{
			1216135, 1216136, 1216137, 1216138, 1216139, 1216140, 1216141, 1216142, 1216143,
			1216144, 1216145, 1216146, 1216147, 1216148, 1216149, 1216150, 1216151, 1216152,
			1216153, 1216154, 1216155, 1216156, 1216157, 1216158, 1216159, 1216160, 1216161,
			1216162, 1216163, 1216164, 1216165, 1216166, 1216167, 1216168, 1216169, 1216170,
		}

		ctx = context.Background()
		db  = NewDatabaseMock(mc)
		cli = NewParserMock(mc)
	)

	// setting up responses from database
	db.IsCachedMock.Expect(ctx, tournamentKey).Return(true, tournamentID, nil)
	db.AddSessionMock.Expect(ctx, tournamentID, chatID).Return(nil)
	db.FetchQuestionsMock.Expect(ctx, tournamentID).Return(questions, nil)

	gameMock, _ := New(cli, db)
	q, err := gameMock.Begin(ctx, tournamentKey, chatID)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(q, questions) {
		t.Fatal(err)
	}
}

func TestGame_FindTournament(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	var (
		tournamentKey = "olivje22_u"
		ctx           = context.Background()
		cli           = NewParserMock(mc)
		db            = NewDatabaseMock(mc)

		tournament = models.Tournament{TextId: tournamentKey}
	)

	cli.TournamentMock.Expect(ctx, tournamentKey).Return(tournament, nil)

	gameMock, _ := New(cli, db)

	trmnt, err := gameMock.FindTournament(ctx, tournamentKey)
	if err != nil {
		t.Fatal(err)
	}

	if trmnt.TextId != tournament.TextId {
		t.Fatal(err)
	}

}

func TestGame_FindTournamentFailure(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	var (
		tournamentKey = "incorrectKey"
		ctx           = context.Background()
		cli           = NewParserMock(mc)
		db            = NewDatabaseMock(mc)

		tournament = models.Tournament{}
	)

	cli.TournamentMock.Expect(ctx, tournamentKey).Return(tournament, ErrChgkNotFound)

	gameMock, _ := New(cli, db)

	_, err := gameMock.FindTournament(ctx, tournamentKey)
	assert.Error(t, err, ErrChgkNotFound)

}

func TestGame_FetchQuestion(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	var (
		questiodID int64 = 1216135
		ctx              = context.Background()
		cli              = NewParserMock(mc)
		db               = NewDatabaseMock(mc)

		question = models.Question{}
	)

	db.FetchQuestionMock.Expect(ctx, questiodID).Return(question, nil)

	gameMock, _ := New(cli, db)

	_, err := gameMock.FetchQuestion(ctx, questiodID)
	assert.Nil(t, err)
}

func TestGame_SubmitGuess(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	var (
		questiodID    int64 = 1216135
		chatID        int64 = 1
		username            = "durov"
		guessWithTypo       = "фитоняфка"
		ctx                 = context.Background()
		cli                 = NewParserMock(mc)
		db                  = NewDatabaseMock(mc)

		question = models.Question{Answer: "фитоняшка"}
	)

	db.FetchQuestionMock.Expect(ctx, questiodID).Return(question, nil)
	db.AddScoreMock.Return(nil)

	gameMock, _ := New(cli, db)

	isCorrect, err := gameMock.SubmitGuess(ctx, chatID, questiodID, username, guessWithTypo)

	assert.Equal(t, isCorrect, true)
	assert.Nil(t, err)

}

func TestGame_SubmitGuessIncorrect(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	var (
		questiodID     int64 = 1216135
		chatID         int64 = 1
		username             = "durov"
		guessIncorrect       = "дрыщ"
		ctx                  = context.Background()
		cli                  = NewParserMock(mc)
		db                   = NewDatabaseMock(mc)

		question = models.Question{Answer: "фитоняшка"}
	)

	db.FetchQuestionMock.Expect(ctx, questiodID).Return(question, nil)
	db.AddScoreMock.Return(nil)

	gameMock, _ := New(cli, db)

	isCorrect, err := gameMock.SubmitGuess(ctx, chatID, questiodID, username, guessIncorrect)

	assert.Equal(t, isCorrect, false)
	assert.Nil(t, err)

}

func TestGame_SubmitGuessFailure(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	var (
		qIncorrectID int64 = 0
		chatID       int64 = 1
		username           = "durov"
		guess              = "фитоняшнка"
		ctx                = context.Background()
		cli                = NewParserMock(mc)
		db                 = NewDatabaseMock(mc)

		question = models.Question{}
	)

	db.FetchQuestionMock.Expect(ctx, qIncorrectID).Return(question, ErrNotFound)

	gameMock, _ := New(cli, db)

	_, err := gameMock.SubmitGuess(ctx, chatID, qIncorrectID, username, guess)

	assert.Error(t, err, ErrNotFound)

}

func TestGame_FetchScore(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	var (
		chatID   int64 = 1
		score    int64 = 0
		username       = "durov"
		ctx            = context.Background()
		cli            = NewParserMock(mc)
		db             = NewDatabaseMock(mc)
		users          = []models.User{{Name: username, Score: score}}
	)

	db.FetchScoreMock.Expect(ctx, chatID).Return(users, nil)

	gameMock, _ := New(cli, db)

	u, err := gameMock.FetchScore(ctx, chatID)
	assert.Equal(t, u, users)
	assert.Nil(t, err)

}

func TestGame_FetchScoreFailure(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	var (
		chatID int64 = 1
		ctx          = context.Background()
		cli          = NewParserMock(mc)
		db           = NewDatabaseMock(mc)
	)

	db.FetchScoreMock.Expect(ctx, chatID).Return(nil, ErrNotFound)

	gameMock, _ := New(cli, db)

	u, err := gameMock.FetchScore(ctx, chatID)
	assert.Nil(t, u)
	assert.Error(t, err, ErrNotFound)
}

func TestGame_FetchSessions(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	var (
		chatID   int64 = 1
		ctx            = context.Background()
		cli            = NewParserMock(mc)
		db             = NewDatabaseMock(mc)
		timeN          = time.Now()
		sessions       = []models.Session{{
			Tournament: "olivje22_u",
			Title:      "Кубок оливье — 2022",
			IsActive:   true,
			Created:    timeN,
		}}
	)

	db.FetchSessionsMock.Expect(ctx, chatID).Return(sessions, nil)

	gameMock, _ := New(cli, db)

	s, err := gameMock.FetchSessions(ctx, chatID)
	assert.Equal(t, s, sessions)
	assert.Nil(t, err)
}

func TestGame_FetchSessionsFailure(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	var (
		chatID int64 = 1
		ctx          = context.Background()
		cli          = NewParserMock(mc)
		db           = NewDatabaseMock(mc)
	)

	db.FetchSessionsMock.Expect(ctx, chatID).Return(nil, ErrNotFound)

	gameMock, _ := New(cli, db)

	s, err := gameMock.FetchSessions(ctx, chatID)
	assert.Error(t, err)
	assert.Nil(t, s)
}

func TestGame_FinishSession(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	var (
		chatID int64 = 1
		ctx          = context.Background()
		cli          = NewParserMock(mc)
		db           = NewDatabaseMock(mc)
	)

	db.FinishSessionMock.Expect(ctx, chatID).Return(nil)

	gameMock, _ := New(cli, db)

	err := gameMock.FinishSession(ctx, chatID)
	assert.Nil(t, err)
}

func TestGame_FinishSessionFailure(t *testing.T) {

	mc := minimock.NewController(t)
	defer mc.Finish()

	var (
		chatID int64 = 1
		ctx          = context.Background()
		cli          = NewParserMock(mc)
		db           = NewDatabaseMock(mc)
	)

	db.FinishSessionMock.Expect(ctx, chatID).Return(ErrNotFound)

	gameMock, _ := New(cli, db)

	err := gameMock.FinishSession(ctx, chatID)
	assert.Error(t, err, ErrNotFound)
}
