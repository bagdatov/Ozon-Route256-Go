package postgres

import (
	"context"

	"gitlab.ozon.dev/bagdatov/homework-2/game-master/usecase"
)

// FinishSession marks session as finished.
func (d *db) FinishSession(ctx context.Context, chatID int64) error {

	update := `UPDATE sessions
			   SET is_active = false
			   WHERE chat_id = $1`

	tag, err := d.pgx.Exec(ctx, update, chatID)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return usecase.ErrNotStarted
	}

	return err
}
