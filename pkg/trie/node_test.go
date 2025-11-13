package trie_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unstoppablemango/x12/pkg/trie"
)

var _ = Describe("Node", func() {
	It("should work", func() {
		// root := trie.Node[string]{{
		// 	label: "a",
		// 	target: trie.Node[string]{{
		// 		label:  "b",
		// 		target: trie.Node[string]{},
		// 	}},
		// }}
		root := trie.New[string, string]()

		node, found := root.Lookup("ab")

		Expect(found).To(BeTrue())
		Expect(node).To(Equal(""))
	})
})
