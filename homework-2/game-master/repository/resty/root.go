package resty

import (
	"context"
	"encoding/xml"
	"fmt"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/models"

	"gitlab.ozon.dev/bagdatov/homework-2/game-master/usecase"
)

// Root is parsing root of the site with all tournaments.
func (c *cli) Root(ctx context.Context) (models.Tournament, error) {

	t := models.Tournament{}

	resp, err := c.resty.R().SetContext(ctx).Get(root)
	if err != nil {
		return t, err
	}

	if resp.IsError() {
		return t, fmt.Errorf(
			"%w, status: %d, response: %s",
			usecase.ErrChgkUnavailable, resp.StatusCode(), resp.Body(),
		)
	}

	if err := xml.Unmarshal(resp.Body(), &t); err != nil {
		return t, err
	}

	if len(t.Tours) == 0 {
		return t, usecase.ErrChgkNotFound
	}

	return t, nil
}
