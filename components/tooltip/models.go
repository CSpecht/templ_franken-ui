package tooltip

import (
	"context"
	"fmt"
	"io"

	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/animation"
	"github.com/cspecht/templ_franken-ui/components/component"
)

type tooltip struct {
	component.C
	ratio     float32
	content   templ.Component
	title     string
	position  Position
	offset    int
	animation animation.Animation
	duration  int
	delay     int
	cls       string
	container string
}

func NewTooltip(title string, content templ.Component) *tooltip {
	return &tooltip{
		title:     title,
		content:   content,
		position:  TopCenter,
		offset:    0,
		animation: animation.ScaleUp,
		duration:  100,
		delay:     0,
		cls:       "uk-active",
		container: "body",
	}
}

func (t *tooltip) WithPosition(position Position) *tooltip {
	t.position = position
	return t
}

func (t *tooltip) WithOffset(offset int) *tooltip {
	t.offset = offset
	return t
}

func (t *tooltip) WithAnimation(animation animation.Animation) *tooltip {
	t.animation = animation
	return t
}

func (t *tooltip) WithDuration(duration int) *tooltip {
	t.duration = duration
	return t
}

func (t *tooltip) WithDelay(delay int) *tooltip {
	t.delay = delay
	return t
}

// Define a target container via a selector to specify where the tooltip should be appended in the DOM.
func (t *tooltip) WithTargetContainer(container string) *tooltip {
	t.container = container
	return t
}
func (t *tooltip) getOptions() string {

	return "title: " + t.title +
		"; pos: " + string(t.position) +
		"; offset: " + fmt.Sprintf("%d", t.offset) +
		"; animation: " + t.animation.String() +
		"; duration: " + fmt.Sprintf("%d", t.duration) +
		"; delay: " + fmt.Sprintf("%d", t.delay) +
		"; cls: " + t.cls +
		"; container: " + t.container
}

// in case if you want to add this to any component
func (t *tooltip) MergeAttributes(attr templ.Attributes) templ.Attributes {
	if attr == nil {
		attr = templ.Attributes{}
	}
	attr["uk-tooltip"] = t.getOptions()
	return attr
}

func (t *tooltip) Render(ctx context.Context, w io.Writer) error {

	return t.component().Render(ctx, w)
}
