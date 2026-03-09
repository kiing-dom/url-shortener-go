package repository

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/kiing-dom/url-shortener-go/internal/model"
	"github.com/redis/go-redis/v9"
)

type RedisURLRepository struct {
	client *redis.Client
}

func NewRedisURLRepository(addr string) *RedisURLRepository {
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return &RedisURLRepository{client: client}
}

func (r *RedisURLRepository) Save(url, code string) error {
	ctx := context.Background()

	exists, err := r.client.Exists(ctx, "url:"+code).Result()
	if err != nil {
		return err
	}

	if exists > 0 {
		return errors.New("URL already exists")
	}

	entry := &model.URLEntry{
		Code:        code,
		OriginalURL: url,
		CreatedAt:   time.Now(),
		Clicks:      0,
	}

	data, err := json.Marshal(entry)
	if err != nil {
		return err
	}

	r.client.Set(ctx, "url:"+code, data, 0)
	r.client.Set(ctx, "reverse:"+url, code, 0)
	return nil
}
