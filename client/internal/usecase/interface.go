package usecase

import (
	"context"
	"thumbs/client/config"
)

// ThumbUseCase is a model layer interface
type ThumbUseCase interface {
	ExecSync(ctx context.Context, id []string, flags config.Cli, errChan chan<- string)
	ExecAsync(ctx context.Context, id []string, flags config.Cli, errChan chan<- string)
}

// ThumbFile is an interface for files
type ThumbFile interface {
	Create(ctx context.Context, id string, data []byte) (string, error)
}
