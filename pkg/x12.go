package x12

import (
	"context"

	"github.com/unstoppablemango/x12/pkg/app"
)

type (
	Handler     = app.Handler[*State]
	HandlerFunc = app.HandlerFunc[*State]
	Option      = app.Option[*State]
	Add         = app.Add[*State]
)

type State struct {
	ctx  context.Context
	path string
}

func NewState(ctx context.Context, path string) *State {
	return &State{ctx, path}
}

func (r *State) Context() context.Context {
	return r.ctx
}

func (r *State) Path() app.Path {
	return app.Path(r.path)
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

func HandleFunc(path string, handler func(*State)) Option {
	return Handle(path, HandlerFunc(handler))
}
