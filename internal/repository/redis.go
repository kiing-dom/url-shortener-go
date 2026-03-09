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
		return errors.New("code already exists")
	}

	reverseExists, reverseErr := r.client.Exists(ctx, "reverse:"+url).Result()
	if reverseErr != nil {
		return reverseErr
	}
	if reverseExists > 0 {
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

	if codeToDataErr := r.client.Set(ctx, "url:"+code, data, 0).Err(); codeToDataErr != nil {
		return codeToDataErr
	}
	if urlToCodeErr := r.client.Set(ctx, "reverse:"+url, code, 0).Err(); urlToCodeErr != nil {
		return urlToCodeErr
	}
	return nil
}

func (r *RedisURLRepository) FindByURL(url string) (string, error) {
	ctx := context.Background()
	code, err := r.client.Get(ctx, "reverse:"+url).Result()
	if err == redis.Nil {
		return "", errors.New("URL not found")
	}
	if err != nil {
		return "", err
	}

	return code, nil
}

func (r *RedisURLRepository) FindByCode(code string) (string, error) {
	ctx := context.Background()
	data, err := r.client.Get(ctx, "url:"+code).Result()

	if err == redis.Nil {
		return "", errors.New("Code not found")
	}
	if err != nil {
		return "", err
	}

	var entry model.URLEntry
	if err := json.Unmarshal([]byte(data), &entry); err != nil {
		return "", err
	}
	return entry.OriginalURL, nil
}

func (r *RedisURLRepository) FindEntryByCode(code string) (*model.URLEntry, error) {
	ctx := context.Background()
	data, err := r.client.Get(ctx, "url:"+code).Result()

	if err == redis.Nil {
		return nil, errors.New("Code not found")
	}
	if err != nil {
		return nil, err
	}

	var entry *model.URLEntry
	if err := json.Unmarshal([]byte(data), &entry); err != nil {
		return nil, err
	}
	return entry, nil
}

func (r *RedisURLRepository) IncrementClicks(code string) error {
	ctx := context.Background()
	data, entryGetErr := r.client.Get(ctx, "url:"+code).Result()
	if entryGetErr != nil {
		return entryGetErr
	}

	var entry *model.URLEntry
	if unmarshallErr := json.Unmarshal([]byte(data), &entry); unmarshallErr != nil {
		return unmarshallErr
	}

	entry.Clicks++
	marshallData, marshallErr := json.Marshal(entry)
	if marshallErr != nil {
		return marshallErr
	}

	if updateErr := r.client.Set(ctx, "url:"+code, marshallData, 0).Err(); updateErr != nil {
		return updateErr
	}
	return nil
}
