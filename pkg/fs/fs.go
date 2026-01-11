package fs

import (
	"io/fs"

	"github.com/unstoppablemango/x12/pkg/app"
)

type (
	File = fs.File
	FS   = fs.FS

	App    = app.App[Request]
	Option = app.Option[Request]
)

type Request struct {
	path string
}

func (r Request) Path() app.Path {
	return app.Path(r.path)
}

type filesystem struct {
	trie *App
}

func (fs *filesystem) Open(name string) (File, error) {
	panic("not implemented")
}

func New(options ...Option) FS {
	return &filesystem{trie: app.New(options...)}
}
