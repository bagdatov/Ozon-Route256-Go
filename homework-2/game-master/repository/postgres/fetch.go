package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/models"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/usecase"
)

// FetchQuestion returns question from database.
func (d *db) FetchQuestion(ctx context.Context, questionID int64) (models.Question, error) {
	q := models.Question{}

	query := `SELECT 
			  id, num, tour_id, question, answer, authors, comments
			  FROM questions
			  WHERE id = $1`

	err := d.pgx.QueryRow(ctx, query, questionID).Scan(
		&q.QuestionId,
		&q.Number,
		&q.ParentId,
		&q.Question,
		&q.Answer,
		&q.Authors,
		&q.Comments,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return q, usecase.ErrNotFound
		}
	}

	return q, err
}

// FetchQuestions returns all question IDs that are related to tournament from database.
func (d *db) FetchQuestions(ctx context.Context, tournamentID int64) ([]int64, error) {

	query := `SELECT q.id
			  FROM questions q
			  INNER JOIN tours t ON q.tour_id = t.id
			  WHERE t.tournament_id = $1
			  ORDER BY q.id 
			  ASC`

	rows, err := d.pgx.Query(ctx, query, tournamentID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, usecase.ErrNotFound
		}
		return nil, err
	}
	defer rows.Close()

	var questionIDs []int64

	for rows.Next() {

		var id int64

		if err := rows.Scan(&id); err != nil {
			return nil, err
		}

		questionIDs = append(questionIDs, id)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return questionIDs, nil
}

// FetchScore returns users of this chat with their scores from database.
// Users are present in database only if they answered previously to questions.
func (d *db) FetchScore(ctx context.Context, chatID int64) ([]models.User, error) {

	query := `SELECT username, COUNT(is_correct)
			  FROM scores
			  WHERE chat_id = $1
			  AND is_correct = true
			  GROUP BY username`

	rows, err := d.pgx.Query(ctx, query, chatID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, usecase.ErrNotFound
		}
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var u models.User

		if err := rows.Scan(&u.Name, &u.Score); err != nil {
			return nil, err
		}

		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// FetchSessions returns all sessions that are related to the chat.
func (d *db) FetchSessions(ctx context.Context, chatID int64) ([]models.Session, error) {

	query := `SELECT t.text_id, t.title, s.is_active, s.created 
			  FROM sessions s
			  INNER JOIN tournaments t ON s.tournament_id = t.id
			  WHERE s.chat_id = $1`

	rows, err := d.pgx.Query(ctx, query, chatID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, usecase.ErrNotFound
		}
		return nil, err
	}

	defer rows.Close()

	var sessions []models.Session

	for rows.Next() {

		var s models.Session

		err := rows.Scan(
			&s.Tournament,
			&s.Title,
			&s.IsActive,
			&s.Created,
		)

		if err != nil {
			return nil, err
		}

		sessions = append(sessions, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return sessions, nil
}
