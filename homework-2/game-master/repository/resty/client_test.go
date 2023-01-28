package resty

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/config"
	"gitlab.ozon.dev/bagdatov/homework-2/game-master/usecase"
)

func TestCli_Tournament(t *testing.T) {
	cli, _ := New(&config.Config{
		ChgkBase: config.ChgkBase{
			Url:     "https://db.chgk.info/",
			Timeout: 15 * time.Second,
		},
	})

	var (
		tournamentKey       = "olivje22_u"
		expected      int64 = 72491
	)

	actual, err := cli.Tournament(context.Background(), tournamentKey)
	assert.Nil(t, err)
	assert.Equal(t, expected, actual.ID)
}

func TestCli_TournamentFailure(t *testing.T) {
	cli, _ := New(&config.Config{
		ChgkBase: config.ChgkBase{
			Url:     "https://db.chgk.info/",
			Timeout: 15 * time.Second,
		},
	})

	var (
		tournamentKey = "incorrectKey"
	)

	_, err := cli.Tournament(context.Background(), tournamentKey)
	assert.Equal(t, err, usecase.ErrChgkNotFound)
}

func TestCli_TournamentIncorrectUrl(t *testing.T) {
	cli, _ := New(&config.Config{
		ChgkBase: config.ChgkBase{
			Url:     "https://db.chgk.inf/",
			Timeout: 15 * time.Second,
		},
	})

	var (
		tournamentKey = "incorrectKey"
	)

	_, err := cli.Tournament(context.Background(), tournamentKey)
	assert.Error(t, err, usecase.ErrChgkUnavailable)
}

func TestCli_Tour(t *testing.T) {
	cli, _ := New(&config.Config{
		ChgkBase: config.ChgkBase{
			Url:     "https://db.chgk.info/",
			Timeout: 15 * time.Second,
		},
	})

	var (
		tourKey        = "olivje22_u.1"
		expected int64 = 72492
	)

	tour, err := cli.Tour(context.Background(), tourKey)
	assert.Nil(t, err)
	assert.Equal(t, expected, tour.ID)
}

func TestCli_TourFailure(t *testing.T) {
	cli, _ := New(&config.Config{
		ChgkBase: config.ChgkBase{
			Url:     "https://db.chgk.info/",
			Timeout: 15 * time.Second,
		},
	})

	var (
		tourKey = "incorrectKey"
	)

	_, err := cli.Tour(context.Background(), tourKey)
	assert.Error(t, err, usecase.ErrNotFound)
}
