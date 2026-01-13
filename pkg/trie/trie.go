package trie

import (
	"iter"

	art "github.com/plar/go-adaptive-radix-tree/v2"
)

type Key = art.Key

type Trie[T any] struct{ tree art.Tree }

func New[T any]() *Trie[T] {
	return &Trie[T]{art.New()}
}

func (t *Trie[T]) CopyTo(insert func(Key, T)) {
	for k, v := range t.Iter() {
		insert(k, v)
	}
}

func (t *Trie[T]) Iter() iter.Seq2[Key, T] {
	return func(yield func(Key, T) bool) {
		t.tree.ForEach(func(node art.Node) (cont bool) {
			return yield(node.Key(), node.Value().(T))
		})
	}
}

func (t *Trie[T]) Lookup(key Key) (T, bool) {
	if val, found := t.tree.Search(key); found {
		return val.(T), true
	} else {
		return *new(T), false
	}
}

func (t *Trie[T]) Insert(key Key, value T) {
	_, _ = t.tree.Insert(key, value)
}
