package webapi

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"thumbs/server/internal/entity"
	"thumbs/server/pkg/webclient"
)

func TestGetThumbFromAPI(t *testing.T) {
	ctx := context.Background()

	c := New(webclient.New())

	type TestCase struct {
		name   string
		id     string
		expVal entity.Pic
		expErr error
	}

	cases := []TestCase{{
		name:   "valid",
		id:     "RAItctARwPs",
		expVal: entity.Pic{ID: "RAItctARwPs", Data: []byte("aboba")},
		expErr: nil,
	}, {
		name:   "404",
		id:     "B",
		expVal: entity.Pic{},
		expErr: entity.ErrNotFound,
	},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p, err := c.GetThumbFromAPI(ctx, tc.id)
			require.ErrorIs(t, err, tc.expErr)
			if err != nil {
				require.Nil(t, p.Data)
			} else {
				require.NotNil(t, p.Data)
			}
		})
	}
}
