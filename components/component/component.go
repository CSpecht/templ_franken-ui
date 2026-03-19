package component

import ("github.com/a-h/templ"
tw "github.com/Oudwins/tailwind-merge-go")


type C struct {
	attributes templ.Attributes
	classes    []string

}

func (c *C) SetAttributes(attributes templ.Attributes) *C {
	// better merge the attributes instead of overwriting them, so that you can add attributes in the component and then overwrite them in the usage
	if c.attributes == nil {
		c.attributes = templ.Attributes{}
	}
	for key, value := range attributes {
		c.attributes[key] = value
	}
	return c
}
func (c *C) SetClasses(styles ... string) *C {
	c.classes = append(c.classes, styles...)
	return c
}
func (c *C) GetClasses() []string {
	c.classes = []string{tw.Merge(c.classes...)}
	return c.classes
}
func (c *C) GetAttributes() templ.Attributes {
	return c.attributes
}
// just helper function that returns the merged classes, so that they can be used in the template (using twmerge)
func MergeClasses(styles ... string) string {
	return tw.Merge(styles...)
}

type AttributesStyles struct {
	Attributes templ.Attributes
	Styles []string
}

