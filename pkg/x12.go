package x12

import (
	"context"

	"github.com/unstoppablemango/x12/pkg/app"
)

type (
	Action      = app.Action[*State]
	Handler     = app.Handler[*State]
	HandlerFunc = app.HandlerFunc[*State]
	Path        = app.Path
	Add         = app.Add[*State]
	Option      = app.Option[*State]
)

type State struct {
	ctx context.Context
}

func (r *State) Context() context.Context {
	return r.ctx
}

func New(options ...Option) Handler {
	return app.New(options...)
}

func With(build func(Add)) Option {
	return app.With(build)
}

func Handle(path Path, handler Handler) Option {
	return app.Handle(path, handler)
}
