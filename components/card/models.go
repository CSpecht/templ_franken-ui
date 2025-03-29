package card

import (
	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
)

type card struct{
	component.C
	style string
	title string
	header templ.Component
	body templ.Component
	footer templ.Component

}
func NewCard() *card {
	return &card{}
}
func (c *card) AsDefault() *card {
	c.style = ""
	return c
}
func (c *card) AsPrimary() *card {
	c.style = primary
	return c
}
func (c *card) AsSecondary() *card {
	c.style = secondary
	return c
}
func (c *card) AsDestructive() *card {
	c.style = destructive
	return c
}
func (c *card) SetTitle(title string) *card {
	c.title = title
	return c
}
func (c *card) SetBody(body templ.Component) *card {
	c.body = body
	return c
}
func (c *card) SetFooter(footer templ.Component) *card {
	c.footer = footer
	return c
}
func (c *card) SetHeader(header templ.Component) *card {
	c.header = header
	return c
}