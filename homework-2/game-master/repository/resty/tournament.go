package resty

import (
	"context"
	"encoding/xml"
	"fmt"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/models"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/usecase"
)

// Tournament is parsing site and returns its tournament and related tours details.
func (c *cli) Tournament(ctx context.Context, name string) (models.Tournament, error) {
	t := models.Tournament{}

	resp, err := c.resty.R().SetContext(ctx).Get(tourPrefix + name + postfix)
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

	if t.ID == 0 {
		return t, usecase.ErrChgkNotFound
	}

	return t, nil
}
