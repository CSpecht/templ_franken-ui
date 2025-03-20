package navbar

import (
	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
	"github.com/cspecht/templ_franken-ui/components/icon"
)

type navbar struct {
	component.C
	variant variant
	items   []templ.Component
}

type navbarEntry struct {
	component.C
	name           string
	item           templ.Component
	active         bool
	link           templ.SafeURL
	linkAttributes templ.Attributes
	icon           *icon.Icon
}

func NewNavbar() *navbar {
	return &navbar{
		variant: Default,
	}
}

func (*navbar) AddNavbarEntry(name string) *navbarEntry {
	return &navbarEntry{name: name, active: false}
}

func (n *navbar) AddHeader(name string) *navbar {
	n.items = append(n.items, n.createHeader(name))
	return n
}

func (n *navbar) AddDivider() *navbar {
	n.items = append(n.items, n.createDivider())
	return n
}

func (ne *navbarEntry) IsActive() *navbarEntry {
	ne.active = true
	return ne
}

func (ne *navbarEntry) SetLink(link templ.SafeURL, attributes templ.Attributes) *navbarEntry {
	ne.link = link
	ne.linkAttributes = attributes
	return ne
}

func (ne *navbarEntry) SetIcon(icon *icon.Icon, attributes templ.Attributes) *navbarEntry {
	ne.icon = icon
	ne.linkAttributes = attributes
	return ne
}

func (ne *navbarEntry) SetSubComponent(item templ.Component) *navbarEntry {
	ne.item = item
	return ne
}

func (n *navbar) AddItem(entry *navbarEntry) *navbar {
	n.items = append(n.items, entry.Render())
	return n
}
