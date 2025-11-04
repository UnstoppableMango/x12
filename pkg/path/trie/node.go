package trie

import "strings"

type Edge struct {
	Label  string
	Target Node
}

type Node []Edge

func (n Node) IsLeaf() bool {
	return len(n) == 0
}

func (n Node) Lookup(path string) (Node, bool) {
	node, found := n, 0
	for node != nil && !node.IsLeaf() && found < len(path) {
		node, found = node.next(path, found)
	}

	return node, node != nil && node.IsLeaf() && found == len(path)
}

func (n Node) next(path string, found int) (Node, int) {
	for _, edge := range n {
		if strings.HasPrefix(path[found:], edge.Label) {
			return edge.Target, found + len(edge.Label)
		}
	}

	return nil, 0
}
