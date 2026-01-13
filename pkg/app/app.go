package app

import (
	"fmt"

	"github.com/unmango/go/option"
	"github.com/unstoppablemango/x12/pkg/trie"
)

type Request interface {
	Path() Path
}

type Path = trie.Key

type HandlerFunc[T Request] func(T)

func (handle HandlerFunc[T]) Handle(req T) {
	handle(req)
}

type Trie[T Request] interface {
	CopyTo(func(Path, Handler[T]))
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
	return From(trie.New[Handler[T]](), options...)
}

func From[T Request](trie Trie[T], options ...Option[T]) *App[T] {
	app := &App[T]{trie: trie}
	option.ApplyAll(app, options)
	return app
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
