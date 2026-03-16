package spinner

import (
	"context"
	"fmt"
	"io"

	"github.com/cspecht/templ_franken-ui/components/component"
)

type spinner struct {
	component.C
	ratio float32
}

func NewSpinner() *spinner {
	return &spinner{
		ratio: 1.0,
	}
}

func (s *spinner) WithRatio(ratio float32) *spinner {
	s.ratio = ratio
	return s
}
func (s *spinner) getRatio() string {
	if s.ratio != 1.0 {
		return "ratio: " + fmt.Sprintf("%f", s.ratio)
	} else {
		return ""
	}
}
func (s *spinner) Render(ctx context.Context, w io.Writer) error {

	return s.component().Render(ctx, w)
}
