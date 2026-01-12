package path12

import "github.com/unstoppablemango/x12/pkg/app"

type (
	Path  = app.Path
	State = app.State

	Handler[T State] = app.Handler[Piped[T]]
	Option[T State]  = app.Option[Piped[T]]
)

type Piped[T State] interface {
	State
	Base() T
}

type piped[T State] struct {
	base T
	path Path
}

func (p *piped[T]) Base() T {
	return p.base
}

func (p *piped[T]) Path() Path {
	return p.path
}

func Pipe[T State](base T, fns ...func(Path) Path) Piped[T] {
	p := &piped[T]{
		base: base,
		path: base.Path(),
	}

	for _, fn := range fns {
		p.path = fn(p.path)
	}

	return p
}
