package command

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
)

type command struct {
	component.C
	key       string
	comp templ.Component
}

func NewCommand(key string) *command {
	return &command{key: key}
}
func (c *command) SetComponent(component templ.Component) *command {
	c.comp = component
	return c
}

func (c *command) Render(ctx context.Context, w io.Writer) error {

	return c.component().Render(ctx, w)
}
