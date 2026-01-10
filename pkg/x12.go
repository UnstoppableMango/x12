package x12

import (
	"context"

	"github.com/unstoppablemango/x12/pkg/app"
)

type (
	Action       = app.Action[*Request]
	ErrorHandler = app.ErrorHandler
	Handler      = app.Handler[*Request]
	HandlerFunc  = app.HandlerFunc[*Request]
	Path         = app.Path
	Add          = app.Add[*Request]
	Option       = app.Option[*Request]
	Builder      func(Add)
)

type Request struct {
	ctx context.Context
}

func NewRequest(ctx context.Context) *Request {
	return &Request{ctx: ctx}
}

func (r *Request) Context() context.Context {
	return r.ctx
}

func New(options ...Option) Handler {
	return app.New(options...)
}

func With(build Builder) Option {
	return app.With(build)
}

func Handle(path Path, handler Handler) Option {
	return app.Handle(path, handler)
}
