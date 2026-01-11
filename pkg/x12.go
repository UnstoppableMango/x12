package x12

import (
	"bytes"
	"context"
	"io"
	"sync"

	"github.com/unstoppablemango/x12/pkg/app"
)

type Request interface {
	app.State
	io.Reader
	io.Writer

	Context() context.Context
	Err(error)
}

type (
	App         = app.App[Request]
	Handler     = app.Handler[Request]
	HandlerFunc = app.HandlerFunc[Request]
	Option      = app.Option[Request]
	Add         = app.Add[Request]
)

func New(options ...Option) *App {
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

func NotFound(handler func(string)) Option {
	return app.NotFound[Request](func(p app.Path) {
		handler(string(p))
	})
}

type request struct {
	mu   sync.RWMutex
	path string
	ctx  context.Context
	buf  bytes.Buffer
}

func Req(path string, options ...func(*request)) Request {
	return &request{path: path}
}

func (r *request) Context() context.Context {
	return r.ctx
}

func (r *request) Path() app.Path {
	return app.Path(r.path)
}

func (r *request) Read(p []byte) (int, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.buf.Read(p)
}

func (r *request) Write(p []byte) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.buf.Write(p)
}

func (r *request) Err(err error) {
	if _, err := r.Write([]byte(err.Error())); err != nil {
		panic(err)
	}
}

func WithContext(ctx context.Context) func(*request) {
	return func(req *request) {
		req.ctx = ctx
	}
}
