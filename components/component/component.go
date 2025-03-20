package component

import "github.com/a-h/templ"


type C struct {
	attributes templ.Attributes
	classes    []string
}

func (c *C) SetAttributes(attributes templ.Attributes) *C {
	c.attributes = attributes
	return c
}
func (c *C) AddClasses(styles ... string) *C {
	c.classes = append(c.classes, styles...)
	return c
}
func (c *C) GetClasses() []string {
	return c.classes
}
func (c *C) GetAttributes() templ.Attributes {
	return c.attributes
}
