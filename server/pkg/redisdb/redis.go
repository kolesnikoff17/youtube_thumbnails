package redisdb

import (
	"context"
	"github.com/go-redis/redis/v8"
)

// Conn is a connection to redis db
type Conn struct {
	Rdb *redis.Client
}

// New is a constructor for Conn
func New(address string) (*Conn, error) {
	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
	_, err := rdb.Get(ctx, "key").Result()
	if err != redis.Nil {
		return nil, err
	}
	return &Conn{
		Rdb: rdb,
	}, nil
}
