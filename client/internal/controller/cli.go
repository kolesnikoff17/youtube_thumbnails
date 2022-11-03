package controller

import (
	"context"
	"flag"
	"fmt"
	"log"
	"thumbs/client/config"
	"thumbs/client/internal/usecase"
)

// Cli is a controller for client app
type Cli struct {
	uc usecase.ThumbUseCase

	flags config.Cli
}

// New is a constructor for Cli
func New(uc usecase.ThumbUseCase, flags config.Cli) *Cli {
	return &Cli{
		uc:    uc,
		flags: flags,
	}
}

// Exec is an entrypoint of business-logic
func (c *Cli) Exec(ctx context.Context, done chan<- struct{}) {
	ids := c.parseArgs()
	if ids == nil {
		log.Fatal("No ids to download")
	}

	errChan := make(chan string, 1)
	defer close(errChan)

	go func() {
		for msg := range errChan {
			fmt.Println(msg)
		}
		done <- struct{}{}
	}()

	if c.flags.Async {
		c.uc.ExecAsync(ctx, ids, c.flags, errChan)
	} else {
		c.uc.ExecSync(ctx, ids, c.flags, errChan)
	}
}

func (c *Cli) parseArgs() []string {
	return flag.Args()
}
