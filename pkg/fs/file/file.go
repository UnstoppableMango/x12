package file

import (
	"io/fs"

	"github.com/unstoppablemango/x12/pkg/app"
)

type Request interface {
	app.Request

	Err(error)
}

type (
	Path    = app.Path
	Handler = app.Handler[Request]
)

type Op int

const (
	OpClose Op = iota
	OpRead
	OpStat
)

type request struct {
	Handler
	op Op
}

// Close implements [fs.File].
func (req *request) Close() error {
	panic("unimplemented")
}

// Read implements [fs.File].
func (req *request) Read([]byte) (int, error) {
	panic("unimplemented")
}

// Stat implements [fs.File].
func (req *request) Stat() (fs.FileInfo, error) {
	panic("unimplemented")
}
