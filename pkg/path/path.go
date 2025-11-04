package path

import x12 "github.com/unstoppablemango/x12/pkg"

type Router interface{}

type Builder interface {
	Build() Router
}

func Execute(path x12.Path) error {
	panic("unimplemented")
}
