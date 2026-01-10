package http

import (
	"context"
	"net"
	"net/http"

	"github.com/unstoppablemango/x12/pkg/app"
)

type (
	App         = app.App[*State]
	Option      = app.Option[*State]
	Handler     = app.Handler[*State]
	HandlerFunc = app.HandlerFunc[*State]

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

func (state *State) Path() app.Path {
	return app.Path(state.Req.RequestURI)
}

func Handle(path string, handler Handler) {
	DefaultServeMux.Handle(path, handler)
}

func HandleFunc(path string, handler func(*State) error) {
	DefaultServeMux.HandleFunc(path, handler)
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
	handler Handler
	error   func(error)
}

func adapt(handler Handler) *adapter {
	return &adapter{
		handler: handler,
		error: func(err error) {
			panic(err)
		},
	}
}

func (a *adapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s := &State{Res: w, Req: r}
	a.handler.Handle(s)
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
	mux.app.Insert(app.Path(path), handler)
}

func (mux *ServeMux) HandleFunc(path string, handler func(*State) error) {
	mux.Handle(path, HandlerFunc(handler))
}

func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
	s := &State{Res: w, Req: r}
	if err := mux.app.Handle(s); err != nil {
		mux.error(err)
	}
}

func (mux *ServeMux) SetErrorHandler(handler func(error)) {
	mux.error = handler
}
