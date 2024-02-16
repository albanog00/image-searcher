package db

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	Client *redis.Client
}

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.TODO()
)

func NewRedisClient() (*RedisClient, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     ":6379",
		Password: "",
		DB:       0,
	})

	if err := client.Ping(Ctx).Err(); err != nil {
		return nil, errors.New("Can't reach RedisClient")
	}

	return &RedisClient{
		Client: client,
	}, nil
}
