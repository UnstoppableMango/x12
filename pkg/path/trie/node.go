package trie

import (
	"strings"

	"github.com/unstoppablemango/x12/pkg/path"
)

type Key interface {
	string | []byte
}

type edge[K Key, T any] struct {
	label  K
	target node[K, T]
}

type node[K Key, T any] struct {
	edges []edge[K, T]
	value T
}

func (n *node[K, T]) Insert(path K, value T) {
	panic("not implemented")
}

func (n *node[K, T]) IsLeaf() bool {
	return n != nil && len(n.edges) == 0
}

func (n node[K, T]) Lookup(path K) (T, bool) {
	node, found := n, 0
	for !node.IsLeaf() && found < len(path) {
		node, found = node.next(path, found)
	}

	return node.value, node.IsLeaf() && found == len(path)
}

func (n node[K, T]) next(path K, found int) (node[K, T], int) {
	for _, edge := range n.edges {
		if hasPrefix(path[found:], edge.label) {
			return edge.target, found + len(edge.label)
		}
	}

	return node[K, T]{}, 0
}

func hasPrefix[K Key](path K, prefix K) bool {
	return strings.HasPrefix(string(path), string(prefix))
}

func New[K Key, T any]() path.Trie[K, T] {
	return &node[K, T]{}
}

func Node[K Key, T any](v T) path.Trie[K, T] {
	return &node[K, T]{value: v}
}
