package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgconn"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/models"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/usecase"
)

// AddTournament using one transaction to fill database with parsed tournament, tours, and questions.
func (d *db) AddTournament(ctx context.Context, trmnt models.Tournament, tours []models.Tour) (err error) {

	tx, err := d.pgx.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback(ctx)
		} else {
			err = tx.Commit(ctx)
		}
	}()

	insertTournament := `INSERT INTO tournaments 
			  			(id, text_id, title, tours_num, questions_num, created)
			  			VALUES ($1, $2, $3, $4, $5, $6)`

	_, err = tx.Exec(ctx, insertTournament,
		trmnt.ID,
		trmnt.TextId,
		trmnt.Title,
		trmnt.ChildrenNum,
		trmnt.QuestionsNum,
		trmnt.CreatedAt,
	)
	if err != nil {
		return err
	}

	insertTour := `INSERT INTO tours
				   (id, text_id, tournament_id, title, editors, questions_num)
				   VALUES ($1, $2, $3, $4, $5, $6)`

	for _, t := range tours {

		_, err = tx.Exec(ctx, insertTour,
			t.ID,
			t.TextId,
			t.ParentId,
			t.Title,
			t.Editors,
			t.QuestionsNum,
		)
		if err != nil {
			return err
		}
	}

	insertQuestion := `INSERT INTO questions
					  (id, num, tour_id, question, answer, authors, comments)
					  VALUES ($1, $2, $3, $4, $5, $6, $7)`

	for _, t := range tours {
		for _, q := range t.Questions {

			_, err = tx.Exec(ctx, insertQuestion,
				q.QuestionId,
				q.Number,
				q.ParentId,
				q.Question,
				q.Answer,
				q.Authors,
				q.Comments,
			)
			if err != nil {
				return err
			}
		}
	}

	return err
}

// AddSession starts session within database.
// Returns error if tournament in this chat already played.
func (d *db) AddSession(ctx context.Context, tournamentID, chatID int64) error {

	insert := `INSERT INTO sessions
			   (chat_id, tournament_id, is_active, created)
			   VALUES ($1, $2, true, $3)`
	_, err := d.pgx.Exec(ctx, insert, chatID, tournamentID, time.Now())
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == errUniqueViolation {
			err = usecase.ErrStarted
		}
	}

	return err
}

// AddScore marks question as answered for this user.
// Returns error if there was an answer from user already.
func (d *db) AddScore(ctx context.Context, chatID, questionID int64, user string, isCorrect bool) error {

	insert := `INSERT INTO scores
			   (chat_id, username, question_id, is_correct, created)
			   VALUES ($1, $2, $3, $4, $5)`

	_, err := d.pgx.Exec(ctx, insert, chatID, user, questionID, isCorrect, time.Now())
	if err != nil {
		var pgErr *pgconn.PgError

		if errors.As(err, &pgErr) && pgErr.Code == errUniqueViolation {
			err = usecase.ErrAnswered
		}
	}
	return err
}
