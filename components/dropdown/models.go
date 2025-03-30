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
}

func NewDropdown(triggerComponent templ.Component, nav nav.Nav) *dropdown {
	nav.AddClasses("uk-dropdown-nav")
	d := &dropdown{
		triggerComponent: triggerComponent,
		navComponent:     nav,
	}
	return d
}
