package usecase

import (
	"context"
	"fmt"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/models"
	"math/rand"
	"reflect"
	"strings"
	"sync"
	"time"
)

// game implements Game interface
type game struct {
	parser Parser
	db     Database
}

// New is a constructor for game
func New(cli Parser, db Database) (*game, error) {
	if isNil(cli) {
		return nil, fmt.Errorf("parser is nil")
	}

	if isNil(db) {
		return nil, fmt.Errorf("database is nil")
	}

	return &game{
		parser: cli,
		db:     db,
	}, nil
}

// isNil is using reflect
// package to determine rather input is nil
func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}

// Begin start new session and parsing questions from chgk website by using tournament key.
// Tournament key must not contain dots.
func (g *game) Begin(ctx context.Context, tournamentKey string, chatID int64) ([]int64, error) {
	if strings.ContainsRune(tournamentKey, '.') {
		return nil, ErrChgkBadRequst
	}

	cached, tID, err := g.db.IsCached(ctx, tournamentKey)
	if err != nil {
		return nil, fmt.Errorf("IsCashed: %w", err)
	}

	if !cached {

		tournament, err := g.parser.Tournament(ctx, tournamentKey)
		if err != nil {
			return nil, fmt.Errorf("Parser.Tournament: %w", err)
		}

		tours, err := g.parseTours(ctx, tournament)
		if err != nil {
			return nil, err
		}

		if err := g.db.AddTournament(ctx, tournament, tours); err != nil {
			return nil, fmt.Errorf("AddTournament: %w", err)
		}
		tID = tournament.ID
	}

	if err := g.db.AddSession(ctx, tID, chatID); err != nil {
		return nil, fmt.Errorf("AddSession: %w", err)
	}

	qIDs, err := g.db.FetchQuestions(ctx, tID)
	if err != nil {
		return nil, fmt.Errorf("FetchQuestions: %w", err)
	}

	return qIDs, nil
}

// parseTours is parsing tours asynchronously
func (g *game) parseTours(ctx context.Context, tournament models.Tournament) ([]models.Tour, error) {

	tours := make([]models.Tour, 0, len(tournament.Tours))

	var (
		err error
		wg  sync.WaitGroup
		mu  sync.Mutex
	)

	// за счет асинхронности сэкономил в среднем 100ms при 3 турах
	// флаг --race не выявил ошибок
	for i := 0; i < len(tournament.Tours); i++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			name := tournament.Tours[id].TextId
			tour, errTour := g.parser.Tour(ctx, name)
			if errTour != nil {
				err = fmt.Errorf("Parser.Tour: %w, tourname: %s", errTour, name)
				return
			}

			mu.Lock()
			tours = append(tours, tour)
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	if err != nil {
		return nil, err
	}

	return tours, nil
}

// FindTournament returns details of tournament. It is parsing website directly and not using cache.
func (g *game) FindTournament(ctx context.Context, key string) (models.Tournament, error) {

	t, err := g.parser.Tournament(ctx, key)
	if err != nil {
		return t, fmt.Errorf("Parser.Tournament: %w", err)
	}

	return t, nil
}

// FetchQuestion returns question details from cache.
func (g *game) FetchQuestion(ctx context.Context, questionID int64) (models.Question, error) {

	q, err := g.db.FetchQuestion(ctx, questionID)
	if err != nil {
		return q, fmt.Errorf("FetchQuestion: %w", err)
	}

	return q, nil
}

// SubmitGuess inserts answer from player into database.
// It is using Jaro Distance to determine rather answer is correct, so guesses with minor typos should be counted.
func (g *game) SubmitGuess(ctx context.Context, chatID, questionID int64, username, guess string) (bool, error) {
	q, err := g.db.FetchQuestion(ctx, questionID)
	if err != nil {
		return false, fmt.Errorf("FetchQuestion: %w", err)
	}

	var isCorrect bool

	// TODO:: revise minimum similarity
	if JaroDistance(q.Answer, guess) >= minSimilarity {
		isCorrect = true
	}

	// not saving answer itself only correctness
	if err := g.db.AddScore(ctx, chatID, questionID, username, isCorrect); err != nil {
		return isCorrect, fmt.Errorf("AddScore: %w", err)
	}

	return isCorrect, nil
}

// FetchScore returns leaderboard of chat. Users are sorted by decreasing order.
func (g *game) FetchScore(ctx context.Context, chatID int64) ([]models.User, error) {
	users, err := g.db.FetchScore(ctx, chatID)
	if err != nil {
		return nil, fmt.Errorf("FetchScore: %w", err)
	}

	return users, nil
}

// FetchSessions returns all sessions of chat.
func (g *game) FetchSessions(ctx context.Context, chatID int64) ([]models.Session, error) {
	sessions, err := g.db.FetchSessions(ctx, chatID)
	if err != nil {
		return nil, fmt.Errorf("FetchSessions: %w", err)
	}

	return sessions, nil
}

// FinishSession finishes current session of chat.
func (g *game) FinishSession(ctx context.Context, chatID int64) error {
	err := g.db.FinishSession(ctx, chatID)
	if err != nil {
		return fmt.Errorf("FinishSession: %w", err)
	}

	return nil
}

// RandomTournaments returns 5 random tournaments from root of chgk site.
// It is not using cache and parses site directly.
func (g *game) RandomTournaments(ctx context.Context) ([]models.TourI, error) {
	rand.Seed(time.Now().Unix())

	root, err := g.parser.Root(ctx)
	if err != nil {
		return nil, fmt.Errorf("RandomTournaments: %w", err)
	}

	tours := make([]models.TourI, 0, 5)

	for len(tours) != 5 {
		i := rand.Intn(len(root.Tours))
		tours = append(tours, root.Tours[i])
	}

	return tours, nil
}
