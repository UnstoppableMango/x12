package http

import (
	"context"
	"net"
	"net/http"

	"github.com/unstoppablemango/x12/pkg/app"
)

type (
	App    = app.App[*State]
	Action = app.Action[*State]
	Option = app.Option[*State]

	Handler        = http.Handler
	HandlerFunc    = http.HandlerFunc
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

func Handle(path string, handler Handler) {
	DefaultServeMux.Handle(path, handler)
}

func HandleFunc(path string, handler func(w ResponseWriter, r *Request)) {
	DefaultServeMux.HandleFunc(path, handler)
}

func Serve(l net.Listener, handler Handler) error {
	srv := &Server{Handler: handler}
	return srv.Serve(l)
}

func ServeTLS(l net.Listener, handler Handler, certFile, keyFile string) error {
	srv := &Server{Handler: handler}
	return srv.ServeTLS(l, certFile, keyFile)
}

var defaultServeMux = ServeMux{
	app: app.New[*State](),
	error: func(err error) {
		panic(err)
	},
}

var DefaultServeMux = &defaultServeMux

type ServeMux struct {
	app   *App
	error func(error)
}

// NewServeMux creates an empty ServeMux that panics on errors.
func NewServeMux() *ServeMux {
	mux := defaultServeMux
	return &mux
}

func (mux *ServeMux) Handle(path string, handler Handler) {
	mux.app.Insert(app.Path(path), func(s *State) error {
		handler.ServeHTTP(s.Res, s.Req)
		return nil
	})
}

func (mux *ServeMux) HandleFunc(path string, handler func(w ResponseWriter, r *Request)) {
	mux.Handle(path, HandlerFunc(handler))
}

func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
	s := &State{Res: w, Req: r}
	p := app.Path(r.RequestURI)
	if err := mux.app.Handle(p, s); err != nil {
		mux.error(err)
	}
}

func (mux *ServeMux) SetErrorHandler(handler func(error)) {
	mux.error = handler
}
