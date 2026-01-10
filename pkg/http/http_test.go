package http_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unstoppablemango/x12/pkg/http"
)

var _ = Describe("Http", func() {
	Describe("NewServeMux", func() {
		It("should create a new ServeMux", func() {
			mux := http.NewServeMux()

			Expect(mux).NotTo(BeNil())
			Expect(mux).NotTo(BeEquivalentTo(http.DefaultServeMux))
		})
	})
})
