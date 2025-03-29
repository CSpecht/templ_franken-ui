package container

import (
	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
)

type container struct {
	component.C
	size   string
	content templ.Component
}

func NewContainer(content templ.Component) *container {
	return &container{
		content: content,
	}
}
func (c *container) IsDefault() *container {
	c.size = ""
	return c
}
func (c *container) IsXSmall() *container {
	c.size =xsmall
	return c
}
func (c *container) IsSmall() *container {
	c.size = small
	return c
}
func (c *container) IsLarge() *container {
	c.size = large
	return c
}
func (c *container) IsXLarge() *container {
	c.size = xlarge
	return c
}
func (c *container) IsExpanded() *container {
	c.size = expanded
	return c
}
func (c *container) SetContent(content templ.Component) *container {
	c.content = content
	return c
}