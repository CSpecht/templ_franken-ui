package command

import (
	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
)

type command struct {
	component.C
	key       string
	component templ.Component
}

func NewCommand(key string) *command {
	return &command{key: key}
}
func (c *command) SetComponent(component templ.Component) *command {
	c.component = component
	return c
}
