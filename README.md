# x12

A micro-framework for building invocable strings in Go.

Probably the same as something that exists, I didn't search very hard beforehand.
A thin wrapper around <https://github.com/plar/go-adaptive-radix-tree> where every value is a [Handler](./pkg/app/app.go).
Package [app](./pkg/app/app.go) holds the core generic abstractions.
Package [x12](./pkg/x12.go) implements a rudimentary request handler using `app`, with a [net/http](https://pkg.go.dev/net/http) style API.
Package [http](./pkg/http/http.go) implements a thin adapter for [net/http](https://pkg.go.dev/net/http) to demonstrate how the request model maps across domains.

```go
app := x12.New(x12.HandleFunc("/user", func(req x12.Request) error {
	if _, err := io.WriteString(req, "Hello World!"); err != nil {
		return err
	}
	return nil
}))

req := x12.Req("/user")
app.Handle(req)

data, _ := io.ReadAll(req)
Expect(string(data)).To(Equal("Hello World!"))
```
