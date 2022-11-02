package usecase

import (
	"context"
	"thumbs/server/internal/entity"
)

// ThumbUseCase is a model layer interface
type ThumbUseCase interface {
	GetThumb(ctx context.Context, id string) (entity.Pic, error)
}

// ThumbRepo is a repository layer interface
type ThumbRepo interface {
	Get(ctx context.Context, id string) (entity.Pic, error)
	Put(ctx context.Context, pic entity.Pic) error
}

// ThumbWebAPI is an interface for third-party API
type ThumbWebAPI interface {
	GetThumbFromAPI(ctx context.Context, id string) (entity.Pic, error)
}
