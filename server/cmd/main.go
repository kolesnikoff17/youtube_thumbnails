package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "thumbs/proto"
	"thumbs/server/config"
	"thumbs/server/internal/controller/grpcserver"
	"thumbs/server/internal/usecase"
	"thumbs/server/internal/usecase/repository"
	"thumbs/server/internal/usecase/webapi"
	"thumbs/server/pkg/logger"
	"thumbs/server/pkg/redisdb"
	"thumbs/server/pkg/webclient"
	"time"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("failed to get cfg: %s", err)
	}
	l, err := logger.New(cfg.Log.Lvl)
	if err != nil {
		log.Fatalf("failed to configure logger: %s", err)
	}
	db, err := redisdb.New(cfg.Db.Host + ":" + cfg.Db.Port)
	if err != nil {
		l.Fatalf("failed to connect to db: %s", err)
	}
	c := webclient.New(webclient.MaxConn(cfg.Client.MaxConn),
		webclient.IdleTimeout(time.Duration(cfg.Client.IdleTO)*time.Second))
	uc := usecase.New(repository.New(db), webapi.New(c))
	s := grpcserver.New(uc, l)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Grpc.Port))
	if err != nil {
		l.Fatalf("failed to listen port: %s", err)
	}
	server := grpc.NewServer()
	pb.RegisterThumbServer(server, s)
	err = server.Serve(lis)
	if err != nil {
		l.Fatalf("failed to serve: %s", err)
	}
}
