package grpcserver

import (
  "context"
  "errors"
  "google.golang.org/grpc/codes"
  "google.golang.org/grpc/status"
  pb "thumbs/proto"
  "thumbs/server/internal/entity"
  "thumbs/server/internal/usecase"
  "thumbs/server/pkg/logger"
)

// Server implements gRPC server
type Server struct {
  uc usecase.ThumbUseCase
  l  logger.Interface
  pb.UnimplementedThumbServer
}

// New is a constructor for Server
func New(uc usecase.ThumbUseCase, l logger.Interface) *Server {
  return &Server{
    uc: uc,
    l:  l,
  }
}

// GetThumbnail return picture as a byte stream
func (s *Server) GetThumbnail(ctx context.Context, req *pb.ThumbRequest) (*pb.ThumbResponse, error) {
  p, err := s.uc.GetThumb(ctx, req.GetId(), req.GetUpdate())
  switch {
  case errors.Is(err, entity.ErrNotFound):
    return nil, status.Error(codes.NotFound, "No such id")
  case err != nil:
    s.l.Warnf("err %s with id %s", err, req.GetId())
    return nil, status.Error(codes.Internal, "Internal error")
  }
  return &pb.ThumbResponse{Thumb: p.Data}, nil
}
