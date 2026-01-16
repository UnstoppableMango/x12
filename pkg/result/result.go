package result

import (
	"github.com/unmango/go/either/result"
	"github.com/unstoppablemango/x12/pkg/app"
)

type (
	Path          = app.Path
	Result[T any] = result.Result[T]
)

type Request[T any] interface {
	app.Request

	SetResult(T)
	SetError(error)
}

type Handler[Res any, Req Request[Res]] interface {
	Handle(Req) (Res, error)
}

type HandlerFunc[Res any, Req Request[Res]] func(Req) (Res, error)

func (handle HandlerFunc[Res, Req]) Handle(req Req) (Res, error) {
	return handle(req)
}

func Handle[Res any, Req Request[Res]](path Path, handler Handler[Res, Req]) app.Option[Req] {
	return app.HandleFunc(path, func(req Req) {
		if res, err := handler.Handle(req); err != nil {
			req.SetError(err)
		} else {
			req.SetResult(res)
		}
	})
}

func HandleFunc[Res any, Req Request[Res]](path Path, handler func(Req) (Res, error)) app.Option[Req] {
	return Handle(path, HandlerFunc[Res, Req](handler))
}

func Error[T any](err error) Result[T] {
	return result.Error[T](err)
}

func Ok[T any](val T) Result[T] {
	return result.Ok(val)
}
