package repository

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"thumbs/server/internal/entity"
	"thumbs/server/internal/usecase"
	"thumbs/server/pkg/redisdb"
	"time"
)

// RedisRepo implements usecase.ThumbRepo
type RedisRepo struct {
	c *redisdb.Conn
}

var _ usecase.ThumbRepo = (*RedisRepo)(nil)

// New is a constructor for RedisRepo
func New(c *redisdb.Conn) *RedisRepo {
	return &RedisRepo{
		c: c,
	}
}

// Get returns cached picture from repo by its id
func (r *RedisRepo) Get(ctx context.Context, id string) (entity.Pic, error) {
	val, err := r.c.Rdb.Get(ctx, id).Bytes()
	switch {
	case errors.Is(err, redis.Nil):
		return entity.Pic{}, entity.ErrNotFound
	case err != nil:
		return entity.Pic{}, err
	}
	return entity.Pic{ID: id, Data: val}, nil
}

// Put insert picture in repo
func (r *RedisRepo) Put(ctx context.Context, pic entity.Pic) error {
	res := r.c.Rdb.Set(ctx, pic.ID, pic.Data, time.Hour)
	if res.Err() != nil {
		return res.Err()
	}
	return nil
}
