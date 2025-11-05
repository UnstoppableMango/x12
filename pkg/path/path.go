package path

type Trie[Key, Node any] interface {
	Lookup(path Key) (Node, bool)
}
