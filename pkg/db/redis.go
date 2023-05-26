package db

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type RedisOptions struct {
	Host     string
	Port     string
	DB       int
	Password string
}

func NewRedis(opts *RedisOptions) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     opts.Host + ":" + opts.Port,
		Password: opts.Password, // no password set
		DB:       opts.DB,       // use default DB
	})
	_, err := client.Ping(context.Background()).Result()

	if err != nil {
		return nil, err
	}
	return client, nil
}
