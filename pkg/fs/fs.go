package fs

import (
	"fmt"
	"io/fs"

	"github.com/unstoppablemango/x12/pkg/app"
)

type Request interface {
	fmt.Stringer
	app.State

	Err(error)
	Op() Op
	File(fs.File)
}

type (
	File     = fs.File
	FileInfo = fs.FileInfo
	FS       = fs.FS

	App     = app.App[Request]
	Handler = app.Handler[Request]
	Option  = app.Option[Request]
)

type Op int

const (
	OpOpen Op = iota
	OpClose
)

type request struct {
	op   Op
	path string
	err  error
	file File
}

func (r *request) String() string {
	return r.path
}

func (r request) Path() app.Path {
	return app.Path(r.path)
}

func (r *request) Err(err error) {
	r.err = err
	r.file = nil
}

func (r *request) Op() Op {
	return r.op
}

func (r *request) File(f fs.File) {
	r.file = f
	r.err = nil
}

type filesystem struct{ Handler }

func New(options ...Option) FS {
	return &filesystem{app.New(options...)}
}

func (fs *filesystem) Open(name string) (File, error) {
	req := &request{op: OpOpen, path: name}
	fs.Handle(req)
	return req.file, req.err
}

type Opener struct{ FS }

func (fs Opener) Handle(req Request) {
	if f, err := fs.Open(req.String()); err != nil {
		req.Err(err)
	} else {
		req.File(f)
	}
}

type file struct{ Handler }

func (f file) Handle(req Request) {
	req.File(&fileRequest{f.Handler, req})
}

type fileRequest struct {
	Handler
	req Request
}

// Close implements [fs.File].
func (f *fileRequest) Close() error {
	panic("unimplemented")
}

// Read implements [fs.File].
func (f *fileRequest) Read([]byte) (int, error) {
	panic("unimplemented")
}

// Stat implements [fs.File].
func (f *fileRequest) Stat() (fs.FileInfo, error) {
	panic("unimplemented")
}
