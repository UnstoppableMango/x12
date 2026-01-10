package x12_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	x12 "github.com/unstoppablemango/x12/pkg"
)

var _ = Describe("X12", func() {
	It("should create an app", func() {
		app := x12.New()

		Expect(app).NotTo(BeNil())
	})

	It("should handle a request", func() {
		var flag bool
		app := x12.New(x12.HandleFunc("/", func(s x12.Request) {
			flag = true
		}))

		app.Handle(x12.Request("/"))

		Expect(flag).To(BeTrue())
	})
})
