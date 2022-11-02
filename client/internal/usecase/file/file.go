package file

import (
	"context"
	"os"
	"strconv"
	"strings"
	"thumbs/client/internal/usecase"
	"time"
)

// File implements usecase.ThumbFile
type File struct{}

var _ usecase.ThumbFile = (*File)(nil)

// New is a constructor for File
func New() *File { return &File{} }

// Create make new picture
func (f *File) Create(ctx context.Context, id string, data []byte) error {
	var s strings.Builder
	s.Grow(40)
	s.WriteString("client/thumbs/")
	s.WriteString(id)
	s.WriteString("_")
	s.WriteString(strconv.Itoa(time.Now().Nanosecond()))
	s.WriteString(".jpg")
	file, err := os.Create(s.String())
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
