package app

import (
	"context"
	"fmt"

	"github.com/unmango/go/option"
	"github.com/unstoppablemango/x12/pkg/trie"
)

type (
	Path               = trie.Key
	Action[R Req]      func(R) error
	ErrorHandler       func(Path) error
	HandlerFunc[R Req] func(Path, R) error

	Add[R Req]  = trie.Insert[Path, Action[R]]
	Trie[R Req] = trie.Trie[Path, Action[R]]
)

type Handler[R Req] interface {
	Handle(path Path, req R) error
}

func (handle HandlerFunc[R]) Handle(path Path, req R) error {
	return handle(path, req)
}

type Req interface {
	Context() context.Context
}

type app[R Req] struct {
	trie     Trie[R]
	notFound ErrorHandler
}

type Option[R Req] func(*app[R])

func New[R Req](options ...Option[R]) Handler[R] {
	app := &app[R]{
		trie: trie.New[Action[R]](),
		notFound: func(p Path) error {
			return fmt.Errorf("no route found for path: %s", p)
		},
	}

	option.ApplyAll(app, options)
	return app
}

// Handle implements X12.
func (a *app[R]) Handle(path Path, req R) error {
	if action, found := a.trie.Lookup(path); found {
		return action(req)
	} else {
		return a.notFound(path)
	}
}

func With[R Req](build func(Add[R])) Option[R] {
	return func(a *app[R]) {
		build(a.trie.Insert)
	}
}

func Handle[R Req](path Path, handler Handler[R]) Option[R] {
	return With(func(add Add[R]) {
		add(path, func(req R) error {
			return handler.Handle(path, req)
		})
	})
}
