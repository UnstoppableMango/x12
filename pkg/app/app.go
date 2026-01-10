package app

import (
	"fmt"

	"github.com/unmango/go/option"
	"github.com/unstoppablemango/x12/pkg/trie"
)

type (
	Path                 = trie.Key
	HandlerFunc[T State] func(T)

	Add[T State] = trie.Insert[Path, Handler[T]]
)

type Trie[T State] interface {
	Lookup(Path) (Handler[T], bool)
	Insert(Path, Handler[T])
}

type Handler[T State] interface {
	Handle(state T)
}

func (handle HandlerFunc[T]) Handle(state T) {
	handle(state)
}

type State interface {
	Path() Path
}

type App[T State] struct {
	trie     Trie[T]
	notFound func(Path)
}

type Option[T State] func(*App[T])

func New[T State](options ...Option[T]) *App[T] {
	app := &App[T]{
		trie: trie.New[Handler[T]](),
		notFound: func(p Path) {
			panic(fmt.Sprintf("no route found for path: %s", p))
		},
	}

	option.ApplyAll(app, options)
	return app
}

func (a *App[T]) Handle(state T) {
	path := state.Path()
	if handler, found := a.trie.Lookup(path); found {
		handler.Handle(state)
	} else {
		a.notFound(path)
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
