package result

import "github.com/unstoppablemango/x12/pkg/app"

type Handler[Request app.Request, Response any] interface {
	Handle(Request) (Response, error)
}
