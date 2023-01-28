package postgres

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
)

// IsCached shows if tournament was saved in database previously, and if yes returns its ID.
func (d *db) IsCached(ctx context.Context, tournamentName string) (bool, int64, error) {

	query := `SELECT id
			  FROM tournaments
			  WHERE text_id = $1
			  LIMIT 1`

	var id int64

	err := d.pgx.QueryRow(ctx, query, tournamentName).Scan(&id)
	if errors.Is(err, pgx.ErrNoRows) {
		return false, id, nil
	}
	if err != nil {
		return false, id, err
	}

	return true, id, nil
}
