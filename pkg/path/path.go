package path

type Trie[Key, Node any] interface {
	Insert(path Key, value Node)
	Lookup(path Key) (Node, bool)
}
