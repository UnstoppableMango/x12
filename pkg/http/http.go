package http

import (
	"context"
	"net/http"

	"github.com/unstoppablemango/x12/pkg/app"
)

type (
	Path    string
	Handler = app.Handler[*Request]
	Option  = app.Option[*Request]
)

type Request struct {
	Res http.ResponseWriter
	Req *http.Request
}

func (r *Request) Context() context.Context {
	return r.Req.Context()
}

func Handle(path Path, handler Handler) Option {
	return app.Handle(app.Path(path), handler)
}
