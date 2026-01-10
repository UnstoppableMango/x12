package app

import (
	"context"
	"fmt"

	"github.com/unmango/go/option"
	"github.com/unstoppablemango/x12/pkg/trie"
)

type (
	Path                 = trie.Key
	HandlerFunc[T State] func(T) error

	Add[T State] = trie.Insert[Path, Handler[T]]
)

type Trie[T State] interface {
	Lookup(Path) (Handler[T], bool)
	Insert(Path, Handler[T])
}

type Handler[T State] interface {
	Handle(state T) error
}

func (handle HandlerFunc[T]) Handle(state T) error {
	return handle(state)
}

type State interface {
	Context() context.Context
	Path() Path
}

type App[T State] struct {
	trie     Trie[T]
	notFound func(Path) error
}

type Option[T State] func(*App[T])

func New[T State](options ...Option[T]) *App[T] {
	app := &App[T]{
		trie: trie.New[Handler[T]](),
		notFound: func(p Path) error {
			return fmt.Errorf("no route found for path: %s", p)
		},
	}

	option.ApplyAll(app, options)
	return app
}

func (a *App[T]) Handle(state T) error {
	path := state.Path()
	if handler, found := a.trie.Lookup(path); found {
		return handler.Handle(state)
	} else {
		return a.notFound(path)
	}
}

func (a *App[T]) Insert(path Path, handler Handler[T]) {
	a.trie.Insert(path, handler)
}

func With[T State](build func(Add[T])) Option[T] {
	return func(a *App[T]) {
		build(a.trie.Insert)
	}
}

func Handle[T State](path Path, handler Handler[T]) Option[T] {
	return With(func(add Add[T]) {
		add(path, handler)
	})
}

func HandleFunc[T State](path Path, handler HandlerFunc[T]) Option[T] {
	return Handle(path, handler)
}
