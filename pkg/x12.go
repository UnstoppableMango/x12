package x12

import (
	"github.com/unstoppablemango/x12/pkg/app"
)

type (
	Handler     = app.Handler[Request]
	HandlerFunc = app.HandlerFunc[Request]
	Option      = app.Option[Request]
	Add         = app.Add[Request]
)

type Request string

func (r Request) Path() app.Path {
	return app.Path(r)
}

func New(options ...Option) Handler {
	return app.New(options...)
}

func With(build func(Add)) Option {
	return app.With(build)
}

func Handle(path string, handler Handler) Option {
	return app.Handle(app.Path(path), handler)
}

func HandleFunc(path string, handler func(Request)) Option {
	return Handle(path, HandlerFunc(handler))
}
