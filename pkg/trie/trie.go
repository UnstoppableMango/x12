package trie

import art "github.com/plar/go-adaptive-radix-tree/v2"

type (
	Key = art.Key

	Lookup[K, T any] func(K) (T, bool)
	Insert[K, T any] func(K, T)
)

type Trie[K, T any] interface {
	Lookup(K) (T, bool)
	Insert(K, T)
}

type trie[T any] struct{ art.Tree }

func (t *trie[T]) Lookup(key Key) (T, bool) {
	if val, found := t.Search(key); found {
		return val.(T), true
	} else {
		return *new(T), false
	}
}

func (t *trie[T]) Insert(key Key, value T) {
	t.Insert(key, value)
}

func New[T any]() Trie[Key, T] {
	return &trie[T]{art.New()}
}
