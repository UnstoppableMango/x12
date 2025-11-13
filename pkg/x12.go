package x12

import (
	"context"
)

type X12 interface {
	Handle(ctx Context, path Path) error
}

type Path string

func (p Path) String() string {
	return string(p)
}

type Routes[Key, Node any] interface {
	Insert(path Key, value Node)
	Lookup(path Key) (Node, bool)
}

type Context interface {
	Context() context.Context
}

type Action func(Context) error

type app struct {
	Routes[Path, Action]
}

// Handle implements X12.
func (a *app) Handle(ctx Context, path Path) error {
	panic("unimplemented")
}

func NewApp() X12 {
	return &app{}
}
