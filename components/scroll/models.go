package scroll

import (
	"io"
	"strconv"

	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/components"
	. "maragu.dev/gomponents/html"
)
// would do same like with the C component
type compi struct {
	Classes   Classes
	HTMLAttrs Group
}

func (c *compi) AddClass(classes ...string) *compi {
	for _, class := range classes {
		c.Classes[class] = true
	}
	return c
}
func (c *compi) AddHTMLAttr(attrs ...Node) *compi {
	for _, attr := range attrs {
		c.HTMLAttrs = append(c.HTMLAttrs, attr)
	}
	return c
}

type scroll struct {
	compi
	content Node
	target  string
	offset  int
}

func NewScroll(target string, content Node) *scroll {

	return &scroll{
		content: content,
		target:  target,
	}
}
func (s *scroll) WithOffset(offset int) *scroll {
	s.offset = offset
	return s
}
// When we implement the Render function we can simple pass it to any other Node, and also simple render on http server
func (s *scroll) Render(w io.Writer) error {
	return A(
		Href(s.target),
		//If(s.offset > 0,Data("uk-scroll", "offset: "+strconv.Itoa(s.offset)), // does not support else

		func() Node {
			if s.offset > 0 {
				return Data("uk-scroll", "offset: "+strconv.Itoa(s.offset))
			}
			return Data("uk-scroll", "")
		}(),	
	
		Group(s.HTMLAttrs),
		s.Classes,
		s.content,
	).Render(w)
}
