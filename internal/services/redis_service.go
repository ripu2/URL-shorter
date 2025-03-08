package services

import (
	"context"
	"errors"
	"time"

	RedisClient "example.com/url-shorter/config/redis"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func SetValueInCache(key, value string) error {
	err := RedisClient.RedisClient.Set(ctx, key, value, 24*time.Hour).Err()
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func GetValueFromCache(key string) (string, error) {
	longURL, err := RedisClient.RedisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return longURL, nil
}
