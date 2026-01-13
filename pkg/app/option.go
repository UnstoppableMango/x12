package app

import "iter"

type (
	Insert[T Request] func(Path, Handler[T])
	Option[T Request] func(*App[T])
)

func Builder[T Request](build func(Insert[T])) Option[T] {
	return func(a *App[T]) {
		build(a.trie.Insert)
	}
}

func Handle[T Request](path Path, handler Handler[T]) Option[T] {
	return Builder(func(add Insert[T]) {
		add(path, handler)
	})
}

func HandleAll[T Request](handlers iter.Seq2[Path, Handler[T]]) Option[T] {
	return Builder(func(add Insert[T]) {
		for p, h := range handlers {
			add(p, h)
		}
	})
}

func HandleFunc[T Request](path Path, handler HandlerFunc[T]) Option[T] {
	return Handle(path, handler)
}

func NotFound[T Request](handler func(Path)) Option[T] {
	return func(app *App[T]) {
		app.notFound = handler
	}
}
