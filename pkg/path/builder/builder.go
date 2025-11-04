package builder

import "github.com/unstoppablemango/x12/pkg/path"

type builder struct{}

// Build implements path.Builder.
func (b *builder) Build() path.Router {
	panic("unimplemented")
}

func New() path.Builder {
	return &builder{}
}
