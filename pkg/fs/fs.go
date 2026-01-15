package fs

import (
	"fmt"
	"io/fs"

	"github.com/unmango/go/option"
	"github.com/unstoppablemango/x12/pkg/app"
)

type Request interface {
	fmt.Stringer
	app.Request
}

type (
	Handler = app.Handler[Request]
	Opener  = app.Handler[OpenRequest]
	Option  = func(*FS)
)

type OpenRequest interface {
	Request

	File(fs.File)
	Err(error)
}

type FS struct{ Opener }

func New(options ...Option) fs.FS {
	fs := &FS{}
	option.ApplyAll(fs, options)
	return fs
}

func (f *FS) Open(name string) (fs.File, error) {
	if !fs.ValidPath(name) {
		return nil, f.pathErr(name, fs.ErrInvalid)
	}

	req := &opener{name: name}

	f.Handle(req)
	if req.err != nil {
		return nil, req.err
	}
	if req.file != nil {
		return req.file, nil
	}

	return nil, f.pathErr(name, fs.ErrNotExist)
}

func (o *FS) pathErr(name string, err error) error {
	return &fs.PathError{
		Op:   "open",
		Path: name,
		Err:  err,
	}
}

type opener struct {
	name string
	err  error
	file fs.File
}

// Err implements [OpenRequest].
func (o *opener) Err(err error) {
	o.file = nil
	o.err = err
}

// File implements [OpenRequest].
func (o *opener) File(file fs.File) {
	o.err = nil
	o.file = file
}

// Path implements [Request].
func (o *opener) Path() app.Path {
	return app.Path(o.name)
}

// String implements [Request].
func (o *opener) String() string {
	return o.name
}
