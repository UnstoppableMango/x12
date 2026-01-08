package x12

import (
	"context"
	"fmt"

	"github.com/unmango/go/option"
	"github.com/unstoppablemango/x12/pkg/trie"
)

type X12 interface {
	Handle(ctx Context, path Path) error
}

type Path string

func (p Path) String() string {
	return string(p)
}

type Routes[Key, Node any] interface {
	Lookup(path Key) (Node, bool)
}

type Context interface {
	Context() context.Context
}

type (
	Action       func(Context) error
	ErrorHandler func(Path) error
)

type app struct {
	Routes[Path, Action]
	notFound ErrorHandler
}

type Option func(*app)

// Handle implements X12.
func (a *app) Handle(ctx Context, path Path) error {
	if action, found := a.Lookup(path); found {
		return action(ctx)
	} else {
		return a.notFound(path)
	}
}

func New(options ...Option) X12 {
	app := &app{
		Routes: trie.New[Path, Action](),
		notFound: func(p Path) error {
			return fmt.Errorf("no route found for path: %s", p)
		},
	}

	option.ApplyAll(app, options)
	return app
}
