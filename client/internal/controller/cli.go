package controller

import (
	"context"
	"flag"
	"fmt"
	"log"
	"thumbs/client/internal/usecase"
)

// Cli is a controller for client app
type Cli struct {
	uc usecase.ThumbUseCase

	upd   bool
	async bool
}

// New is a constructor for Cli
func New(uc usecase.ThumbUseCase, upd bool, async bool) *Cli {
	return &Cli{
		uc:    uc,
		upd:   upd,
		async: async,
	}
}

// Exec is an entrypoint of business-logic
func (c *Cli) Exec(ctx context.Context) {
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
	}()

	if c.async {
		c.uc.ExecAsync(ctx, ids, c.upd, errChan)
	} else {
		c.uc.ExecSync(ctx, ids, c.upd, errChan)
	}
}

func (c *Cli) parseArgs() []string {
	// flag.Parse()
	return flag.Args()
}
