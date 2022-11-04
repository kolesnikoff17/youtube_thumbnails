package repository

import (
	"context"
	"errors"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/require"
	"testing"
	"thumbs/server/internal/entity"
	"thumbs/server/pkg/redisdb"
	"time"
)

func TestGet(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()
	rdb := New(&redisdb.Conn{Rdb: db})

	mock.ExpectGet("A").SetVal("aboba")
	mock.ExpectGet("B").RedisNil()
	mock.ExpectGet("C").SetErr(errors.New("aboba"))

	cases := []struct {
		name   string
		id     string
		expVal entity.Pic
		expErr error
	}{{
		name:   "valid",
		id:     "A",
		expVal: entity.Pic{ID: "A", Data: []byte("aboba")},
		expErr: nil,
	}, {
		name:   "not found",
		id:     "B",
		expVal: entity.Pic{},
		expErr: entity.ErrNotFound,
	}, {
		name:   "db err",
		id:     "C",
		expVal: entity.Pic{},
		expErr: errors.New("aboba"),
	},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p, err := rdb.Get(ctx, tc.id)
			require.Equal(t, tc.expErr, err)
			require.Equal(t, tc.expVal, p)
		})
	}
}

func TestPut(t *testing.T) {
	ctx := context.Background()
	db, mock := redismock.NewClientMock()
	rdb := New(&redisdb.Conn{Rdb: db})

	mock.ExpectSet("A", []byte("aboba"), time.Hour).SetVal("1")
	mock.ExpectSet("B", []byte("aboba"), time.Hour).SetErr(errors.New("aboba"))

	cases := []struct {
		name   string
		pic    entity.Pic
		expErr error
	}{{
		name:   "valid",
		pic:    entity.Pic{ID: "A", Data: []byte("aboba")},
		expErr: nil,
	}, {
		name:   "db err",
		pic:    entity.Pic{ID: "B", Data: []byte("aboba")},
		expErr: errors.New("aboba"),
	},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := rdb.Put(ctx, tc.pic)
			require.Equal(t, tc.expErr, err)
		})
	}
}
