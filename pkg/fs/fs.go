package fs

import "github.com/unstoppablemango/x12/pkg/app"

type (
	App = app.App[Request]
)

type Request struct {
	path string
}

func (r Request) Path() app.Path {
	return app.Path(r.path)
}
