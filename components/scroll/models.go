package scroll

import (
	"context"
	"io"
	"strings"

	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
)

type scroll struct {
	component.C
	target  string
	offset  int
	content templ.Component
}

func NewScroll(targetId string, content templ.Component) *scroll {
	return &scroll{
		target:  targetId,
		content: content,
	}
}
func (s *scroll) getTarget() templ.SafeURL {
	if strings.HasPrefix(s.target, "#") {
		return templ.SafeURL(s.target)
	}
	return templ.SafeURL("#" + s.target)
}

func (s *scroll) WithOffset(offset int) *scroll {
	s.offset = offset
	return s
}
func (s *scroll) getOffset() string {
	if s.offset != 0 {
		return "offset: " + string(s.offset)
	} else {
		return ""
	}
}
func (s *scroll) Render(ctx context.Context, w io.Writer) error {

	return s.component().Render(ctx, w)
}
