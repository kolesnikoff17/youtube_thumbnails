package usecase

import (
	"context"
	"errors"
	"thumbs/server/internal/entity"
)

// Thumb implements ThumbUseCase interface and keeps repo and
// WebAPI interfaces for business-logic implementation
type Thumb struct {
	r ThumbRepo
	c ThumbWebAPI
}

var _ ThumbUseCase = (*Thumb)(nil)

// New is a constructor for Thumb
func New(r ThumbRepo, c ThumbWebAPI) *Thumb {
	return &Thumb{
		r: r,
		c: c,
	}
}

// GetThumb returns picture. It tries to find it in repo, if there is no one it requests it from an API and put in repo.
// Returns entity.ErrNotFound if API returns 404
func (uc *Thumb) GetThumb(ctx context.Context, id string) (entity.Pic, error) {
	p, err := uc.r.Get(ctx, id)
	switch {
	case err == nil:
		return p, nil
	case errors.Is(err, entity.ErrNotFound):
	case err != nil:
		return entity.Pic{}, err
	}
	p, err = uc.c.GetThumbFromAPI(ctx, id)
	switch {
	case errors.Is(err, entity.ErrNotFound):
		return entity.Pic{}, err
	case err != nil:
		return entity.Pic{}, err
	}
	err = uc.r.Put(ctx, p)
	if err != nil {
		return entity.Pic{}, err
	}
	return p, nil
}
