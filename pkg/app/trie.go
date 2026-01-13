package app

import (
	"iter"

	art "github.com/plar/go-adaptive-radix-tree/v2"
)

type key = art.Key

type trie[T any] struct{ tree art.Tree }

func NewTrie[T any]() *trie[T] {
	return &trie[T]{art.New()}
}

func (t *trie[T]) CopyTo(insert func(Path, T)) {
	for k, v := range t.Iter() {
		insert(k, v)
	}
}

func (t *trie[T]) Iter() iter.Seq2[Path, T] {
	return func(yield func(Path, T) bool) {
		t.tree.ForEach(func(node art.Node) bool {
			return yield(node.Key(), node.Value().(T))
		})
	}
}

func (t *trie[T]) Lookup(path Path) (T, bool) {
	if val, found := t.tree.Search(path); found {
		return val.(T), true
	} else {
		return *new(T), false
	}
}

func (t *trie[T]) Insert(path Path, value T) {
	_, _ = t.tree.Insert(path, value)
}
