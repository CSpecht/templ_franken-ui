package dropdown

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/animation"
	"github.com/cspecht/templ_franken-ui/components/component"
	"github.com/cspecht/templ_franken-ui/components/nav"
)

type dropdown struct {
	component.C
	triggerComponent templ.Component
	navComponent     nav.Nav
	mode             string
	animation        animation.Animation
}

func NewDropdown(triggerComponent templ.Component, nav nav.Nav) *dropdown {
	nav.SetClasses("uk-dropdown-nav")
	d := &dropdown{
		triggerComponent: triggerComponent,
		navComponent:     nav,
		animation:        animation.SlideTopSmall,
	}
	return d
}
func (d *dropdown) ModeClickHover() *dropdown {
	d.mode = "mode: click,hover"
	return d
}
func (d *dropdown) ModeClick() *dropdown {
	d.mode = "mode: click"
	return d
}
func (d *dropdown) ModeHover() *dropdown {
	d.mode = "mode: hover"
	return d
}
func (d *dropdown) WithAnimation(animation animation.Animation) *dropdown {
	d.animation = animation
	return d
}
func (d *dropdown) getCustomParameters() string {
	customParams := ""
	if d.mode != "" {
		customParams += d.mode + ";"

	}
	if d.animation != "" {
		customParams += "animation: " + d.animation.String() + ";"
	}
	return customParams
}
func (d *dropdown) Render(ctx context.Context, w io.Writer) error {

	return d.component().Render(ctx, w)
}
