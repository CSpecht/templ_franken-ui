package progress

import (
	"context"
	"io"

	"github.com/cspecht/templ_franken-ui/components/component"
)

type progress struct {
	component.C
	value int
	max   int
}

func NewProgress(value int) *progress {
	return &progress{
		value: value,
		max:   100,
	}
}

func (p *progress) WithMax(max int) *progress {
	p.max = max
	return p
}

func (p *progress) Render(ctx context.Context, w io.Writer) error {

	return p.component().Render(ctx, w)
}
