package http_test

import (
	"fmt"
	"net/url"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/unstoppablemango/x12/pkg/http"
)

type handler struct{ i int }

func (handler) Handle(*http.State) {}

var _ = Describe("Http", func() {
	Describe("ServeMux", func() {
		It("should create a new ServeMux", func() {
			mux := http.NewServeMux()

			Expect(mux).NotTo(BeNil())
			Expect(mux).NotTo(BeIdenticalTo(http.DefaultServeMux))
		})

		DescribeTable("should find handlers",
			func(method, path, wantHandler string) {
				mux := http.FromMap(map[string]http.Handler{
					"/":     &handler{1},
					"/foo/": &handler{2},
					"/foo":  &handler{3},
					"/bar/": &handler{4},
					"//":    &handler{5},
				})

				var r http.Request
				r.Method = method
				r.Host = "example.com"
				r.URL = &url.URL{Path: path}

				h, found := mux.Handler(&http.State{Req: &r})

				Expect(found).To(BeTrue())
				Expect(h).NotTo(BeNil())
				Expect(fmt.Sprintf("%#v", h)).To(MatchRegexp(wantHandler))
			},
			// https://cs.opensource.google/go/go/+/master:src/net/http/server_test.go;l=93-108
			Entry(nil, "GET", "/", "&http_test.handler{i:1}"),
			Entry(nil, "GET", "//", `&http.redirectHandler{url:"/", code:307}`, Pending),
			Entry(nil, "GET", "/foo/../bar/./..//baz", `&http.redirectHandler{url:"/baz", code:307}`, Pending),
			Entry(nil, "GET", "/foo", "&http_test.handler{i:3}"),
			Entry(nil, "GET", "/foo/x", "&http_test.handler{i:2}", Pending),
			Entry(nil, "GET", "/bar/x", "&http_test.handler{i:4}", Pending),
			Entry(nil, "GET", "/bar", `&http.redirectHandler{url:"/bar/", code:307}`, Pending),
			Entry(nil, "CONNECT", "", "(http.HandlerFunc)(.*)", Pending),
			Entry(nil, "CONNECT", "/", "&http_test.handler{i:1}"),
			Entry(nil, "CONNECT", "//", "&http_test.handler{i:1}", Pending),
			Entry(nil, "CONNECT", "//foo", "&http_test.handler{i:5}", Pending),
			Entry(nil, "CONNECT", "/foo/../bar/./..//baz", "&http_test.handler{i:2}", Pending),
			Entry(nil, "CONNECT", "/foo", "&http_test.handler{i:3}"),
			Entry(nil, "CONNECT", "/foo/x", "&http_test.handler{i:2}", Pending),
			Entry(nil, "CONNECT", "/bar/x", "&http_test.handler{i:4}", Pending),
			Entry(nil, "CONNECT", "/bar", `&http.redirectHandler{url:"/bar/", code:307}`, Pending),
		)
	})
})
