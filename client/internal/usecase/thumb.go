package usecase

import (
	"context"
	"fmt"
	"google.golang.org/grpc/status"
	"sync"
	pb "thumbs/proto"
)

// Thumb implements ThumbUseCase
type Thumb struct {
	f ThumbFile
	c pb.ThumbClient
}

var _ ThumbUseCase = (*Thumb)(nil)

// New is a constructor for Thumb
func New(f ThumbFile, c pb.ThumbClient) *Thumb {
	return &Thumb{
		f: f,
		c: c,
	}
}

// ExecSync creates thumbs one by one and type error messages in errChan
func (uc *Thumb) ExecSync(ctx context.Context, id []string, update bool, errChan chan<- string) {
	for _, v := range id {
		uc.getAndCreate(ctx, v, update, errChan)
	}
}

// ExecAsync creates thumbs in async mode and type error messages in errChan
func (uc *Thumb) ExecAsync(ctx context.Context, id []string, update bool, errChan chan<- string) {
	wg := sync.WaitGroup{}
	for _, v := range id {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			uc.getAndCreate(ctx, id, update, errChan)
		}(v)
	}
	wg.Wait()
}

func (uc *Thumb) getAndCreate(ctx context.Context, id string, update bool, errChan chan<- string) {
	r, err := uc.c.GetThumbnail(ctx, &pb.ThumbRequest{Id: id, Update: update})
	if err != nil {
		st := status.Convert(err)
		errChan <- fmt.Sprintf("Err \"%s\" with id %s", st.Message(), id)
		return
	}
	err = uc.f.Create(ctx, id, r.GetThumb())
	if err != nil {
		errChan <- fmt.Sprintf("Err \"Unable to create file\" with id %s", id)
	}
}
