package app

import (
	"fmt"
	"iter"
)

type Path = key // []byte

type Request interface {
	Path() Path
}

type HandlerFunc[T Request] func(T)

func (handle HandlerFunc[T]) Handle(req T) {
	handle(req)
}

type Trie[T Request] interface {
	Iter() iter.Seq2[Path, Handler[T]]
	Lookup(Path) (Handler[T], bool)
	Insert(Path, Handler[T])
}

type Handler[T Request] interface {
	Handle(state T)
}

type App[T Request] struct {
	trie     Trie[T]
	notFound func(Path)
}

func New[T Request](options ...Option[T]) *App[T] {
	return From(NewTrie[Handler[T]](), options...)
}

func (app *App[T]) Handle(req T) {
	path := req.Path()
	if handler, found := app.trie.Lookup(path); found {
		handler.Handle(req)
	} else if app.notFound != nil {
		app.notFound(path)
	} else {
		panic(fmt.Sprintf("no route found for path: %s", path))
	}
}

func (app *App[T]) Lookup(path Path) (Handler[T], bool) {
	return app.trie.Lookup(path)
}

func (app *App[T]) With(options ...Option[T]) *App[T] {
	return From(app.trie, options...)
}
