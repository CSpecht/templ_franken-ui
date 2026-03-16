package close

import (
	"context"
	"io"

	"github.com/cspecht/templ_franken-ui/components/component"
)

type close struct {
	component.C
	button bool
}

func NewButtonClose() *close {
	return &close{button: true}
}
func NewLinkClose() *close {
	return &close{button: false}
}
func (c *close) Render(ctx context.Context, w io.Writer) error {

	return c.component().Render(ctx, w)
}
