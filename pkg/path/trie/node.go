package trie

import "strings"

type Key interface {
	string | []byte
}

type Edge[K Key] struct {
	Label  K
	Target Node[K]
}

type Node[K Key] []Edge[K]

func (n Node[K]) IsLeaf() bool {
	return len(n) == 0
}

func (n Node[K]) Lookup(path K) (Node[K], bool) {
	node, found := n, 0
	for node != nil && !node.IsLeaf() && found < len(path) {
		node, found = node.next(path, found)
	}

	return node, node != nil && node.IsLeaf() && found == len(path)
}

func (n Node[K]) next(path K, found int) (Node[K], int) {
	for _, edge := range n {
		if hasPrefix(path[found:], edge.Label) {
			return edge.Target, found + len(edge.Label)
		}
	}

	return nil, 0
}

func hasPrefix[K Key](path K, prefix K) bool {
	return strings.HasPrefix(string(path), string(prefix))
}
