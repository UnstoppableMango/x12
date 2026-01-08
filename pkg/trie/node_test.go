package trie_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unstoppablemango/x12/pkg/trie"
)

var _ = Describe("Node", func() {
	Describe("Lookup", func() {
		When("the trie is empty", func() {
			It("should return false", func() {
				root := trie.New[string, string]()

				_, found := root.Lookup("ab")

				Expect(found).To(BeFalse())
			})
		})
	})

	Describe("Insert", func() {
		When("inserting a new path", func() {
			It("should be retrievable", func() {
				root := trie.New[string, string]()
				root.Insert("ab", "value")

				value, found := root.Lookup("ab")

				Expect(found).To(BeTrue())
				Expect(value).To(Equal("value"))
			})
		})
	})
})
