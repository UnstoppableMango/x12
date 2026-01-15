package x12

import (
	"bytes"
	"context"
	"io"
	"sync"

	"github.com/unmango/go/fopt"
	"github.com/unstoppablemango/x12/pkg/app"
)

type Request interface {
	app.Request
	io.Reader
	io.Writer

	Context() context.Context
	Err(error)
}

type (
	App    = app.App[Request]
	Option = app.Option[Request]
	Insert = app.Insert[Request]
)

type Handler interface {
	Handle(Request) error
}

type HandlerFunc func(Request) error

func (handle HandlerFunc) Handle(req Request) error {
	return handle(req)
}

func New(options ...Option) *App {
	return app.New(options...)
}

func Builder(build func(Insert)) Option {
	return app.Builder(build)
}

func Handle(path string, handler Handler) Option {
	return app.HandleFunc(app.Path(path), func(req Request) {
		if err := handler.Handle(req); err != nil {
			req.Err(err)
		}
	})
}

func HandleFunc(path string, handler func(Request) error) Option {
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
	err  func(error)
}

func Req(path string, options ...func(*request)) Request {
	req := &request{path: path}
	fopt.ApplyAll(req, options)
	if req.ctx == nil {
		req.ctx = context.Background()
	}

	return req
}

func (req *request) Context() context.Context {
	return req.ctx
}

func (req *request) Path() app.Path {
	return app.Path(req.path)
}

func (req *request) Read(p []byte) (int, error) {
	req.mu.RLock()
	defer req.mu.RUnlock()
	return req.buf.Read(p)
}

func (req *request) Write(p []byte) (int, error) {
	req.mu.Lock()
	defer req.mu.Unlock()
	return req.buf.Write(p)
}

func (req *request) Err(err error) {
	if req.err != nil {
		req.err(err)
	} else {
		panic(err)
	}
}

func WithContext(ctx context.Context) func(*request) {
	return func(req *request) {
		req.ctx = ctx
	}
}

func WithErrorHandler(handler func(error)) func(*request) {
	return func(req *request) {
		req.err = handler
	}
}

func Run(app *App, requests <-chan Request) error {
	return RunContext(context.Background(), app, requests)
}

func RunContext(ctx context.Context, app *App, requests <-chan Request) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case req := <-requests:
			app.Handle(req)
		}
	}
}

func Start(app *App) chan<- Request {
	return StartContext(context.Background(), app)
}

func StartContext(ctx context.Context, app *App) chan<- Request {
	requests := make(chan Request)
	go func() {
		if err := RunContext(ctx, app, requests); err != nil {
			panic(err)
		}
	}()

	return requests
}
