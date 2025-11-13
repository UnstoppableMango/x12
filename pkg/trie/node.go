package trie

import (
	"iter"
	"strings"
)

type Key interface {
	comparable
	string | []byte
}

type Node[K Key, T any] struct {
	edges map[K]Node[K, T]
	value T
}

func (n *Node[K, T]) Insert(path K, value T) {
	panic("not implemented")
}

func (n *Node[K, T]) isLeaf() bool {
	return n != nil && len(n.edges) == 0
}

func (n Node[K, T]) Lookup(path K) (T, bool) {
	node, found := n, 0
	for !node.isLeaf() && found < len(path) {
		node, found = node.next(path, found)
	}

	return node.value, node.isLeaf() && found == len(path)
}

func (n Node[K, T]) next(path K, found int) (Node[K, T], int) {
	for label, edge := range n.edges {
		if hasPrefix(path[found:], label) {
			return edge, found + len(label)
		}
	}

	return Node[K, T]{}, 0
}

func hasPrefix[K Key](path K, prefix K) bool {
	return strings.HasPrefix(string(path), string(prefix))
}

func New[K Key, T any]() *Node[K, T] {
	return &Node[K, T]{}
}

func Single[K Key, T any](v T) *Node[K, T] {
	return &Node[K, T]{value: v}
}

func Iter[K Key, T any](node *Node[K, T]) iter.Seq2[K, T] {
	return nil
}
