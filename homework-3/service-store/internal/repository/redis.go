package repository

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"gitlab.ozon.dev/bagdatov/homework-3/service-store/internal"
	"gitlab.ozon.dev/bagdatov/homework-3/service-store/internal/entity"
	"strconv"
	"time"
)

type cache struct {
	client *redis.Client
	ttl    time.Duration
}

func NewCache(client *redis.Client, ttl time.Duration) *cache {
	return &cache{
		client: client,
		ttl:    ttl,
	}
}

func (c *cache) Get(ctx context.Context, ID uint64) (entity.Item, error) {

	data, err := c.client.Get(ctx, strconv.FormatUint(ID, 10)).Bytes()
	if errors.Is(err, redis.Nil) {
		return entity.Item{}, internal.NotFound
	}
	if err != nil {
		return entity.Item{}, nil
	}

	var item entity.Item

	if err = json.Unmarshal(data, &item); err != nil {
		return entity.Item{}, err
	}

	return item, nil
}

func (c *cache) Set(ctx context.Context, item entity.Item) error {

	data, err := json.Marshal(item)
	if err != nil {
		return err
	}

	return c.client.Set(ctx, strconv.FormatUint(item.ItemID, 10), data, c.ttl).Err()
}

func (c *cache) Delete(ctx context.Context, itemID uint64) error {

	return c.client.Del(ctx, strconv.FormatUint(itemID, 10)).Err()
}
