package usecase

import (
	"context"
	"errors"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
	"thumbs/client/config"
	clientmock "thumbs/client/internal/mocks/client"
	filemock "thumbs/client/internal/mocks/file"
	pb "thumbs/proto"
)

func TestExecAsync(t *testing.T) {
	ctx := context.Background()
	f := filemock.NewThumbFile(t)
	c := clientmock.NewThumbClient(t)
	uc := New(f, c)

	c.On("GetThumbnail", ctx, &pb.ThumbRequest{Id: "A", Update: false}).
		Return(&pb.ThumbResponse{Thumb: []byte("aboba")}, nil)
	f.On("Create", ctx, "A", []byte("aboba")).Return("A", nil)
	c.On("GetThumbnail", ctx, &pb.ThumbRequest{Id: "B", Update: false}).
		Return(nil, status.New(codes.NotFound, "No such id").Err())
	c.On("GetThumbnail", ctx, &pb.ThumbRequest{Id: "C", Update: false}).
		Return(nil, status.New(codes.Internal, "Internal error").Err())
	c.On("GetThumbnail", ctx, &pb.ThumbRequest{Id: "D", Update: false}).
		Return(&pb.ThumbResponse{Thumb: []byte("aboba")}, nil)
	f.On("Create", ctx, "D", []byte("aboba")).Return("", errors.New("aboba"))

	cases := []struct {
		name   string
		id     []string
		flags  config.Cli
		expErr string
	}{{
		name:   "valid",
		id:     []string{"A"},
		flags:  config.Cli{Update: false, Verbose: false},
		expErr: "",
	}, {
		name:   "valid verbose",
		id:     []string{"?v=A"},
		flags:  config.Cli{Update: false, Verbose: true},
		expErr: "Thumb with id A saved into A",
	}, {
		name:   "404",
		id:     []string{`\?v\=B`},
		flags:  config.Cli{Update: false, Verbose: false},
		expErr: `Err "No such id" with id B`,
	}, {
		name:   "500",
		id:     []string{"C"},
		flags:  config.Cli{Update: false, Verbose: false},
		expErr: `Err "Internal error" with id C`,
	}, {
		name:   "failed to create file",
		id:     []string{"D"},
		flags:  config.Cli{Update: false, Verbose: false},
		expErr: `Err "Unable to create file" with id D`,
	},
	}
	ch := make(chan string, 1)

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			uc.ExecAsync(ctx, tc.id, tc.flags, ch)
			if tc.expErr != "" {
				require.Equal(t, tc.expErr, <-ch)
			}
		})
	}
}

func TestExecSync(t *testing.T) {
	ctx := context.Background()
	f := filemock.NewThumbFile(t)
	c := clientmock.NewThumbClient(t)
	uc := New(f, c)

	c.On("GetThumbnail", ctx, &pb.ThumbRequest{Id: "A", Update: false}).
		Return(&pb.ThumbResponse{Thumb: []byte("aboba")}, nil)
	f.On("Create", ctx, "A", []byte("aboba")).Return("A", nil)
	c.On("GetThumbnail", ctx, &pb.ThumbRequest{Id: "B", Update: false}).
		Return(nil, status.New(codes.NotFound, "No such id").Err())
	c.On("GetThumbnail", ctx, &pb.ThumbRequest{Id: "C", Update: false}).
		Return(nil, status.New(codes.Internal, "Internal error").Err())
	c.On("GetThumbnail", ctx, &pb.ThumbRequest{Id: "D", Update: false}).
		Return(&pb.ThumbResponse{Thumb: []byte("aboba")}, nil)
	f.On("Create", ctx, "D", []byte("aboba")).Return("", errors.New("aboba"))

	cases := []struct {
		name   string
		id     []string
		flags  config.Cli
		expErr string
	}{{
		name:   "valid",
		id:     []string{"A"},
		flags:  config.Cli{Update: false, Verbose: false},
		expErr: "",
	}, {
		name:   "valid verbose",
		id:     []string{"?v=A"},
		flags:  config.Cli{Update: false, Verbose: true},
		expErr: "Thumb with id A saved into A",
	}, {
		name:   "404",
		id:     []string{`\?v\=B`},
		flags:  config.Cli{Update: false, Verbose: false},
		expErr: `Err "No such id" with id B`,
	}, {
		name:   "500",
		id:     []string{"C"},
		flags:  config.Cli{Update: false, Verbose: false},
		expErr: `Err "Internal error" with id C`,
	}, {
		name:   "failed to create file",
		id:     []string{"D"},
		flags:  config.Cli{Update: false, Verbose: false},
		expErr: `Err "Unable to create file" with id D`,
	},
	}
	ch := make(chan string, 1)

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			uc.ExecSync(ctx, tc.id, tc.flags, ch)
			if tc.expErr != "" {
				require.Equal(t, tc.expErr, <-ch)
			}
		})
	}
}
