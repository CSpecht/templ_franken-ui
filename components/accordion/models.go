package accordion

import (
	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
)

type accordion struct {
	component.C
	items []*accordionItem
	icon  bool

	customParameters map[string]string
}

type accordionItem struct {
	component.C
	title     string
	text      string
	component templ.Component
	open      bool
	icon      bool
	accordion *accordion
}

func NewAccordion() *accordion {
	return &accordion{icon: true, customParameters: make(map[string]string)}
}
func (a *accordion) DisableIcon() *accordion {
	a.icon = false
	return a
}
func (a *accordion) AddItem(title string, text string) *accordionItem {
	item := &accordionItem{
		title: title,
		text:  text,
		icon:  a.icon,
		accordion: a,
	}
	a.items = append(a.items, item)
	return item
}
func (ai *accordionItem) Accordion() *accordion {
	return ai.accordion
}
func (ai *accordionItem) SetComponent(comp templ.Component) *accordionItem {
	ai.component = comp
	return ai
}
func (ai *accordionItem) IsOpen() *accordionItem {
	ai.open = true
	return ai
}
func (ai *accordionItem) IsClosed() *accordionItem {
	ai.open = false
	return ai
}
func (a *accordion) DisableCollapsing() *accordion {
	a.customParameters["collapsible"] = "false"

	return a
}
func (a *accordion) EnableCollapsing() *accordion {
	a.customParameters["collapsible"] = "true"

	return a
}
func (a *accordion) EnableMultiple() *accordion {
	a.customParameters["multiple"] = "true"
	return a
}
func (a *accordion) DisableMultiple() *accordion {
	a.customParameters["multiple"] = "false"
	return a
}
func (a *accordion) SetCustomParameters(k, v string) *accordion {
	a.customParameters[k] = v
	return a
}
func (a *accordion) getParameters()string{
	params := ""
	for k, v := range a.customParameters {
		params += k + ":" + v + ";"
	}
	return params
}