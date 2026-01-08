package x12_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	x12 "github.com/unstoppablemango/x12/pkg"
)

var _ = Describe("X12", func() {
	It("should work", func() {
		app := x12.New()

		Expect(app).NotTo(BeNil())
	})
})
