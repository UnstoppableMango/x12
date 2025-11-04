package trie_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unstoppablemango/x12/pkg/path/trie"
)

var _ = Describe("Node", func() {
	It("should work", func() {
		root := trie.Node{{
			Label: "a",
			Target: trie.Node{{
				Label:  "b",
				Target: trie.Node{},
			}},
		}}

		node, found := root.Lookup("ab")

		Expect(found).To(BeTrue())
		Expect(node.IsLeaf()).To(BeTrue())
	})
})
