package usecase

import "context"

// ThumbUseCase is a model layer interface
type ThumbUseCase interface {
	ExecSync(ctx context.Context, id []string, update bool, errChan chan<- string)
	ExecAsync(ctx context.Context, id []string, update bool, errChan chan<- string)
}

// ThumbFile is an interface for files
type ThumbFile interface {
	Create(ctx context.Context, id string, data []byte) error
}
