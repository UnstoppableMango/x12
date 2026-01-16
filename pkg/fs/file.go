package fs

import (
	"io/fs"

	"github.com/unstoppablemango/x12/pkg/app"
	"github.com/unstoppablemango/x12/pkg/result"
)

type CloseRequest interface {
	Request
	Err(error)
}

type StatRequest interface {
	result.Request[FileInfo]
}

type ReadRequest interface {
	result.Request[int]
	Data() []byte
}

type (
	Closer = app.Handler[CloseRequest]
	Stater = result.Handler[FileInfo, StatRequest]
	Reader = result.Handler[int, ReadRequest]
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
func (f *file) Read(data []byte) (int, error) {
	return f.read.Handle(&readRequest{
		path: f.name,
		data: data,
	})
}

// Stat implements [fs.File].
func (f *file) Stat() (FileInfo, error) {
	return f.stat.Handle(&statRequest{path: f.name})
}

type closeRequest struct {
	path string
	err  error
}

// Path implements [CloseRequest].
func (c closeRequest) Path() app.Path {
	return app.Path(c.path)
}

// SetError implements [CloseRequest].
func (c *closeRequest) SetError(err error) {
	c.err = err
}

type readRequest struct {
	path string
	data []byte
	res  result.Result[int]
}

// Data implements [ReadRequest].
func (r *readRequest) Data() []byte {
	return r.data
}

// Path implements [ReadRequest].
func (r readRequest) Path() app.Path {
	return app.Path(r.path)
}

// SetError implements [ReadRequest].
func (r *readRequest) SetError(err error) {
	r.res = result.Error[int](err)
}

// SetResult implements [ReadRequest].
func (r *readRequest) SetResult(n int) {
	r.res = result.Ok(n)
}

type statRequest struct {
	path string
	res  result.Result[FileInfo]
}

// Path implements [StatRequest].
func (s *statRequest) Path() app.Path {
	return app.Path(s.path)
}

// SetError implements [StatRequest].
func (s *statRequest) SetError(err error) {
	s.res = result.Error[fs.FileInfo](err)
}

// SetResult implements [StatRequest].
func (s *statRequest) SetResult(info FileInfo) {
	s.res = result.Ok(info)
}
