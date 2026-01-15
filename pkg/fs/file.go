package fs

import (
	"io/fs"

	"github.com/unstoppablemango/x12/pkg/app"
)

type CloseRequest interface {
	Request

	Err(error)
}

type StatRequest interface {
	Request

	Err(error)
	Info(fs.FileInfo)
}

type ReadRequest interface {
	Request

	Data([]byte)
	Err(error)
}

type (
	Closer = app.Handler[CloseRequest]
	Stater = app.Handler[StatRequest]
	Reader = app.Handler[ReadRequest]
)

type file struct {
	name  string
	close Closer
	stat  Stater
	read  Reader
}

func NewFile() fs.File {
	return &file{}
}

// Close implements [fs.File].
func (f *file) Close() error {
	panic("unimplemented")
}

// Read implements [fs.File].
func (f *file) Read([]byte) (int, error) {
	panic("unimplemented")
}

// Stat implements [fs.File].
func (f *file) Stat() (fs.FileInfo, error) {
	panic("unimplemented")
}
