package card

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
)

type card struct {
	component.C
	style            string
	title            string
	titleAttributes  templ.Attributes
	header           templ.Component
	headerAttributes templ.Attributes
	body             templ.Component
	bodyAttributes   templ.Attributes
	footer           templ.Component
	footerAttributes templ.Attributes
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
func (c *card) SetTitle(title string, attributes ...templ.Attributes) *card {
	c.title = title
	if len(attributes) > 0 {
		c.titleAttributes = attributes[0]
	}
	return c
}
func (c *card) SetBody(body templ.Component, attributes ...templ.Attributes) *card {
	c.body = body
	if len(attributes) > 0 {
		c.bodyAttributes = attributes[0]
	}
	return c
}
func (c *card) SetFooter(footer templ.Component, attributes ...templ.Attributes) *card {
	c.footer = footer
	if len(attributes) > 0 {
		c.footerAttributes = attributes[0]
	}
	return c
}
func (c *card) SetHeader(header templ.Component, attributes ...templ.Attributes) *card {
	c.header = header
	if len(attributes) > 0 {
		c.headerAttributes = attributes[0]
	}
	return c
}

func (c *card) Render(ctx context.Context, w io.Writer) error {

	return c.component().Render(ctx, w)
}
