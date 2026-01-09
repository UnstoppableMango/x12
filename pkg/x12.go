package x12

import (
	"context"
	"fmt"

	art "github.com/plar/go-adaptive-radix-tree/v2"
	"github.com/unmango/go/option"
)

type (
	Action[R Req]      func(R) error
	Builder[R Req]     func(Trie[R])
	ErrorHandler       func(Path) error
	HandlerFunc[R Req] func(Path, R) error
	Path               []byte
	Paths[R Req]       func(Path) (Action[R], bool)
	Trie[R Req]        func(Path, Action[R])
)

func (p Path) String() string {
	return string(p)
}

func (p Path) key() art.Key {
	return art.Key(p)
}

type Handler[R Req] interface {
	Handle(path Path, req R) error
}

func (handle HandlerFunc[R]) Handle(path Path, req R) error {
	return handle(path, req)
}

type Req interface {
	Context() context.Context
}

type Request struct {
	ctx context.Context
}

func NewRequest(ctx context.Context) *Request {
	return &Request{ctx: ctx}
}

func (r *Request) Context() context.Context {
	return r.ctx
}

type app[R Req] struct {
	lookup   Paths[R]
	notFound ErrorHandler
}

type Option[R Req] func(*app[R])

func New[R Req](options ...Option[R]) Handler[R] {
	app := &app[R]{
		lookup: func(p Path) (Action[R], bool) {
			return nil, false
		},
		notFound: func(p Path) error {
			return fmt.Errorf("no route found for path: %s", p)
		},
	}

	option.ApplyAll(app, options)
	return app
}

// Handle implements X12.
func (a *app[R]) Handle(path Path, req R) error {
	if action, found := a.lookup(path); found {
		return action(req)
	} else {
		return a.notFound(path)
	}
}

func With[R Req](build Builder[R]) Option[R] {
	return func(a *app[R]) {
		trie := art.New()
		build(func(path Path, action Action[R]) {
			trie.Insert(path.key(), action)
		})

		a.lookup = func(path Path) (Action[R], bool) {
			if val, found := trie.Search(path.key()); found {
				return val.(Action[R]), true
			}
			return nil, false
		}
	}
}

func WithDefault(build Builder[*Request]) Option[*Request] {
	return With(build)
}
