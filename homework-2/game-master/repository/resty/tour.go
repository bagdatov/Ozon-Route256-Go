package resty

import (
	"context"
	"encoding/xml"
	"fmt"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/models"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/usecase"
)

// Tour is parsing site and returns tour detail and related questions.
func (c *cli) Tour(ctx context.Context, name string) (models.Tour, error) {

	t := models.Tour{}

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
