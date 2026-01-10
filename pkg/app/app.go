package app

import (
	"context"
	"fmt"

	"github.com/unmango/go/option"
	"github.com/unstoppablemango/x12/pkg/trie"
)

type (
	Path                 = trie.Key
	Action[T State]      func(T) error
	HandlerFunc[T State] func(Path, T) error

	Add[T State] = trie.Insert[Path, Action[T]]
)

type Trie[T State] interface {
	Lookup(Path) (Action[T], bool)
	Insert(Path, Action[T])
}

type Handler[T State] interface {
	Handle(path Path, state T) error
}

func (handle HandlerFunc[T]) Handle(path Path, state T) error {
	return handle(path, state)
}

type State interface {
	Context() context.Context
}

type App[T State] struct {
	trie     Trie[T]
	notFound func(Path) error
}

type Option[T State] func(*App[T])

func New[T State](options ...Option[T]) *App[T] {
	app := &App[T]{
		trie: trie.New[Action[T]](),
		notFound: func(p Path) error {
			return fmt.Errorf("no route found for path: %s", p)
		},
	}

	option.ApplyAll(app, options)
	return app
}

// Handle implements X12.
func (a *App[T]) Handle(path Path, state T) error {
	if action, found := a.trie.Lookup(path); found {
		return action(state)
	} else {
		return a.notFound(path)
	}
}

func (a *App[T]) Insert(path Path, action Action[T]) {
	a.trie.Insert(path, action)
}

func With[T State](build func(Add[T])) Option[T] {
	return func(a *App[T]) {
		build(a.trie.Insert)
	}
}

func Handle[T State](path Path, handler Handler[T]) Option[T] {
	return HandleFunc(path, handler.Handle)
}

func HandleFunc[T State](path Path, handle HandlerFunc[T]) Option[T] {
	return With(func(add Add[T]) {
		add(path, func(state T) error {
			return handle(path, state)
		})
	})
}
