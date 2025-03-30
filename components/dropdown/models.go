package dropdown

import (
	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
	"github.com/cspecht/templ_franken-ui/components/nav"
)

type dropdown struct {
	component.C
	triggerComponent templ.Component
	navComponent     nav.Nav
	mode             string
}

func NewDropdown(triggerComponent templ.Component, nav nav.Nav) *dropdown {
	nav.AddClasses("uk-dropdown-nav")
	d := &dropdown{
		triggerComponent: triggerComponent,
		navComponent:     nav,
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
func (d *dropdown) getCustomParameters() string {
	customParams := ""
	if d.mode != "" {
		customParams += d.mode + ";"
	}
	return customParams
}