package http

import (
	"context"
	"net"
	"net/http"

	"github.com/unstoppablemango/x12/pkg/app"
)

type (
	Path           string
	App            = app.App[*State]
	Action         = app.Action[*State]
	Handler        = app.Handler[*State]
	Option         = app.Option[*State]
	Server         = http.Server
	ResponseWriter = http.ResponseWriter
	Request        = http.Request
)

type State struct {
	Res http.ResponseWriter
	Req *http.Request
}

func (state *State) Context() context.Context {
	return state.Req.Context()
}

func Handle(path Path, handler Handler) Option {
	return app.Handle(app.Path(path), handler)
}

func Serve(l net.Listener, handler Handler) error {
	srv := &Server{Handler: adapt(handler)}
	return srv.Serve(l)
}

func ServeTLS(l net.Listener, handler Handler, certFile, keyFile string) error {
	srv := &Server{Handler: adapt(handler)}
	return srv.ServeTLS(l, certFile, keyFile)
}

type adapter struct {
	Handler
	error func(error)
}

func adapt(handler Handler) *adapter {
	return &adapter{
		Handler: handler,
		error: func(err error) {
			panic(err)
		},
	}
}

func (a *adapter) ServeHTTP(w ResponseWriter, r *Request) {
	s := &State{Res: w, Req: r}
	p := app.Path(r.RequestURI)
	if err := a.Handle(p, s); err != nil {
		a.error(err)
	}
}

type ServeMux struct{ app *App }

func (mux *ServeMux) Handle(path Path, handler Handler) {
	p := app.Path(path)
	app.Handle(p, handler)
	mux.app.Insert(p, func(s *State) error {
		return handler.Handle(p, s)
	})
}

func NewServeMux() *ServeMux {
	return &ServeMux{app.New[*State]()}
}
