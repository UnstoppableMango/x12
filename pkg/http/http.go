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

func HandleFunc(path string, handler func(*State)) {
	DefaultServeMux.HandleFunc(path, handler)
}

func Serve(l net.Listener, handler Handler) error {
	srv := &Server{Handler: &adapter{handler}}
	return srv.Serve(l)
}

func ServeTLS(l net.Listener, handler Handler, certFile, keyFile string) error {
	srv := &Server{Handler: &adapter{handler}}
	return srv.ServeTLS(l, certFile, keyFile)
}

type adapter struct{ Handler }

func (a *adapter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.Handle(&State{Res: w, Req: r})
}

func New(options ...Option) *App {
	return app.New(options...)
}

var DefaultServeMux = &defaultServeMux

var defaultServeMux = ServeMux{app: New()}

type ServeMux struct {
	app *App
}

func NewServeMux() *ServeMux {
	return &ServeMux{app: New()}
}

func (mux *ServeMux) Handle(path string, handler Handler) {
	mux.app.Insert(app.Path(path), handler)
}

func (mux *ServeMux) HandleFunc(path string, handler func(*State)) {
	mux.Handle(path, HandlerFunc(handler))
}

func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
	mux.app.Handle(&State{Res: w, Req: r})
}
