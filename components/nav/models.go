package nav

import (
	"errors"

	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
	"github.com/cspecht/templ_franken-ui/components/icon"
)

// TODO: Subtitle, Primary Modifier, etc some testing
type item interface {
	Component() templ.Component
}
type nav struct {
	component.C
	variant string
	items   []item
	accordion bool
}

type navItem struct {
	component.C
	name           string
	item           templ.Component
	active         bool
	link           templ.SafeURL
	linkAttributes templ.Attributes
	icon           icon.Icon
	nav            *nav
}

func NewNavigation() *nav {
	return &nav{
		variant: defaultVariant,
	}
}
func (n *nav) AsPrimary() *nav {
	n.variant = primary
	return n
}
func (n *nav) AsSecondary() *nav {
	n.variant = secondary
	return n
}
func (n *nav) AddItem(name string) *navItem {
	e := &navItem{name: name, active: false, nav: n}
	n.items = append(n.items, e)
	return e
}
func (n *nav) AsAccordion()*nav{
	n.accordion = true
	att := n.C.GetAttributes()
	if att == nil {
		att = templ.Attributes{}
	}
	att["data-uk-nav"] = ""
	n.C.SetAttributes(att)
	return n
}
func (n *nav) AllowMultipleOpen() *nav {
	att := n.C.GetAttributes()
	if att == nil {
		att = templ.Attributes{}
	}
	att["data-uk-nav"] = "multiple: true"
	n.C.SetAttributes(att)
	return n
}
func (ne *navItem) Nav() *nav {
	return ne.nav
}

type headerItem struct {
	component.C
	name string
}

func (n *nav) AddHeader(name string) *nav {
	hi := &headerItem{name: name}
	n.items = append(n.items, hi)
	return n
}

type dividerItem struct {
	component.C
}

func (n *nav) AddDivider() *nav {
	n.items = append(n.items, &dividerItem{})
	return n
}

func (ne *navItem) IsActive() *navItem {
	ne.active = true
	return ne
}

func (ne *navItem) SetLink(link templ.SafeURL, attributes ...templ.Attributes) *navItem {
	ne.link = link
	if len(attributes) > 0 {
		ne.linkAttributes = attributes[0]
	}
	return ne
}

func (ne *navItem) SetIcon(icon icon.Icon, attributes ...templ.Attributes) *navItem {
	ne.icon = icon
	if len(attributes) > 0 {
		ne.linkAttributes = attributes[0]
	}

	return ne
}

type navSub struct {
	component.C
	name    string
	mainNav *nav
	parent  *navSub
	items   []item
	parentIcon bool
}

func (n *nav) AddSubNav(name string) *navSub {
	sn := &navSub{name: name, mainNav: n}
	n.items = append(n.items, sn)
	return sn
}
func (ne *navSub) SetParentIcon() *navSub {

	ne.parentIcon = true	
	return ne
}
func (ne *navSub) Nav() *nav {
	return ne.mainNav
}
func (ne *navSub) AddSubNav(name string) *navSub {

	ne2 := &navSub{name: name, parent: ne}
	ne.items = append(ne.items, ne2)
	ne2.mainNav = ne.mainNav
	return ne2
}
func (ne *navSub) Parent() (*navSub, error) {
	if ne.parent == nil {
		return nil, errors.New("no parent")
	}
	return ne.parent, nil
}
func (ne *navSub) AddItem(name string) *navItem {
	e := &navItem{name: name, active: false, nav: ne.Nav()}
	ne.items = append(ne.items, e)
	return e
}
