package listing

import (
	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
)

type list struct {
	component.C
	name    string
	variant variant
	items   []templ.Component
}

type listEntry struct {
	component.C
	name string
	item templ.Component
}

func NewList() *list {
	return &list{
		variant: Default,
	}
}

//	func (l *list) AddListEntry(name string) *listEntry {
//		return &listEntry{name: name}
//	}
// func (*navbar) AddNavbarEntry(name string) *navbarEntry {
// 	return &navbarEntry{name: name, active: false}
// }

func (l *list) AddListEntry(name string) *listEntry {
	return &listEntry{name: name}
}

func (l *list) AddItem(entry *listEntry) *list {
	l.items = append(l.items, entry.Render())
	return l
}

func (l *list) AsStripped() *list {
	l.variant = Stripped
	return l
}

func (l *list) AsPrimary() *list {
	l.variant = Primary
	return l
}

func (l *list) AsSecondary() *list {
	l.variant = Secondary
	return l
}

func (l *list) AsMuted() *list {
	l.variant = Muted
	return l
}

func (l *list) AsDivider() *list {
	l.variant = Divider
	return l
}


