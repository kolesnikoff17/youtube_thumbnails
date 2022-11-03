package usecase

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"testing"
	"thumbs/server/internal/entity"
	repomock "thumbs/server/internal/mocks/redisdb"
	apimock "thumbs/server/internal/mocks/webapi"
)

func TestGetThumb(t *testing.T) {
	ctx := context.Background()
	r := repomock.NewThumbRepo(t)
	c := apimock.NewThumbWebAPI(t)
	uc := New(r, c)

	r.On("Get", ctx, "A").Return(entity.Pic{}, entity.ErrNotFound)
	c.On("GetThumbFromAPI", ctx, "A").
		Return(entity.Pic{ID: "A", Data: []byte("aboba")}, nil)
	r.On("Put", ctx, entity.Pic{ID: "A", Data: []byte("aboba")}).Return(nil)

	r.On("Get", ctx, "B").Return(entity.Pic{ID: "B", Data: []byte("aboba")}, nil)
	c.On("GetThumbFromAPI", ctx, "C").Return(entity.Pic{}, entity.ErrNotFound)
	c.On("GetThumbFromAPI", ctx, "D").Return(entity.Pic{}, errors.New("aboba"))
	c.On("GetThumbFromAPI", ctx, "E").
		Return(entity.Pic{ID: "E", Data: []byte("aboba")}, nil)
	r.On("Put", ctx, entity.Pic{ID: "E", Data: []byte("aboba")}).Return(errors.New("aboba"))
	r.On("Get", ctx, "F").Return(entity.Pic{}, errors.New("aboba"))

	type TestCase struct {
		name   string
		id     string
		expVal entity.Pic
		upd    bool
		expErr error
	}

	cases := []TestCase{{
		name:   "valid no update",
		id:     "A",
		expVal: entity.Pic{ID: "A", Data: []byte("aboba")},
		upd:    false,
		expErr: nil,
	}, {
		name:   "valid update",
		id:     "A",
		expVal: entity.Pic{ID: "A", Data: []byte("aboba")},
		upd:    true,
		expErr: nil,
	}, {
		name:   "valid from cache",
		id:     "B",
		expVal: entity.Pic{ID: "B", Data: []byte("aboba")},
		upd:    false,
		expErr: nil,
	}, {
		name:   "404 from webapi",
		id:     "C",
		expVal: entity.Pic{},
		upd:    true,
		expErr: entity.ErrNotFound,
	}, {
		name:   "500 from webapi",
		id:     "D",
		expVal: entity.Pic{},
		upd:    true,
		expErr: errors.New("aboba"),
	}, {
		name:   "500 from repo put",
		id:     "E",
		expVal: entity.Pic{},
		upd:    true,
		expErr: errors.New("aboba"),
	}, {
		name:   "500 from repo get",
		id:     "F",
		expVal: entity.Pic{},
		upd:    false,
		expErr: errors.New("aboba"),
	},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p, err := uc.GetThumb(ctx, tc.id, tc.upd)
			require.Equal(t, err, tc.expErr)
			require.Equal(t, tc.expVal, p)
		})
	}
}
