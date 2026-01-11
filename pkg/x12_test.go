package x12_test

import (
	"io"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	x12 "github.com/unstoppablemango/x12/pkg"
)

var _ = Describe("X12", func() {
	It("should create an app", func() {
		app := x12.New()

		Expect(app).NotTo(BeNil())
	})

	It("should handle requests", func() {
		var flag bool
		app := x12.New(x12.HandleFunc("/", func(s x12.Request) {
			flag = true
		}))

		app.Handle(x12.Req("/"))

		Expect(flag).To(BeTrue())
	})

	It("should handle requests on different paths", func() {
		var flagA, flagB bool
		app := x12.New(
			x12.HandleFunc("/a", func(s x12.Request) {
				flagA = true
			}),
			x12.HandleFunc("/b", func(s x12.Request) {
				flagB = true
			}),
		)

		app.Handle(x12.Req("/a"))
		Expect(flagA).To(BeTrue())
		Expect(flagB).To(BeFalse())

		app.Handle(x12.Req("/b"))
		Expect(flagB).To(BeTrue())
	})

	It("should count requests per path independently", func() {
		var ca, cb int
		app := x12.New(
			x12.HandleFunc("/a", func(s x12.Request) {
				ca++
			}),
			x12.HandleFunc("/b", func(s x12.Request) {
				cb++
			}),
		)

		app.Handle(x12.Req("/a"))
		Expect(ca).To(Equal(1))
		Expect(cb).To(Equal(0))

		app.Handle(x12.Req("/b"))
		Expect(ca).To(Equal(1))
		Expect(cb).To(Equal(1))

		app.Handle(x12.Req("/b"))
		app.Handle(x12.Req("/b"))
		Expect(ca).To(Equal(1))
		Expect(cb).To(Equal(3))

		app.Handle(x12.Req("/a"))
		Expect(ca).To(Equal(2))
		Expect(cb).To(Equal(3))
	})

	It("should handle IO", func() {
		app := x12.New(x12.HandleFunc("/test", func(s x12.Request) {
			if _, err := io.WriteString(s, "Testing"); err != nil {
				panic(err)
			}
		}))
		req := x12.Req("/test")

		app.Handle(req)

		data, err := io.ReadAll(req)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(data)).To(Equal("Testing"))
	})

	It("should bring receipts for the README", func() {
		app := x12.New(x12.HandleFunc("/user", func(req x12.Request) {
			if _, err := io.WriteString(req, "Hello World!"); err != nil {
				panic(err)
			}
		}))

		req := x12.Req("/user")
		app.Handle(req)
		data, err := io.ReadAll(req)
		Expect(err).NotTo(HaveOccurred())
		Expect(string(data)).To(Equal("Hello World!"))
	})
})
