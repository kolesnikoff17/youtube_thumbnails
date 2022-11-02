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
	ctrl := controller.New(uc, cfg.Cli.Update, cfg.Cli.Async)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		interrupt := make(chan os.Signal, 1)
		defer close(interrupt)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
		<-interrupt
		cancel()
	}()
	ctrl.Exec(ctx)
}
