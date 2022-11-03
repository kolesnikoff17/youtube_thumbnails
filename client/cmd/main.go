package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"os/signal"
	"syscall"
	"thumbs/client/config"
	"thumbs/client/internal/controller"
	"thumbs/client/internal/usecase"
	"thumbs/client/internal/usecase/file"
	pb "thumbs/proto"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to get cfg: %s", err)
	}
	conn, err := grpc.Dial(cfg.Grpc.Host+":"+cfg.Grpc.Port,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect server: %s", err)
	}
	defer conn.Close()
	c := pb.NewThumbClient(conn)
	uc := usecase.New(file.New(), c)
	ctrl := controller.New(uc, cfg.Cli)

	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		ctrl.Exec(ctx, done)
	}()

	select {
	case _ = <-done:
	case _ = <-interrupt:
		cancel()
	}
}
