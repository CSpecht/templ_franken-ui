package listing

import (
	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
)

type list struct {
	component.C
	name    string
	style   string
	variant string
	items   []*listItem
}

type listItem struct {
	component.C
	name string
	item templ.Component
	list *list
}

func NewList() *list {
	return &list{
		variant: defaultVariant,
	}
}

func (l *list) AddItem(name string) *listItem {
	le := &listItem{name: name, list: l}
	l.items = append(l.items, le)
	return le
}
func (le *listItem) AddComponent(c templ.Component) *listItem {
	le.item = c
	return le
}
func (le *listItem) List() *list {
	return le.list
}

func (l *list) AsStriped() *list {
	l.variant = striped
	return l
}
func (l *list) AsBullet() *list {
	l.variant = bullet
	return l
}
func (l *list) AsPrimary() *list {
	l.variant = primary
	return l
}

func (l *list) AsSecondary() *list {
	l.variant = secondary
	return l
}

func (l *list) AsMuted() *list {
	l.variant = muted
	return l
}

func (l *list) AsDivider() *list {
	l.variant = divider
	return l
}

func (l *list) DiscStyle() *list {
	l.style = disc
	return l
}
func (l *list) CircleStyle() *list {
	l.style = circle
	return l
}
func (l *list) SquareStyle() *list {
	l.style = square
	return l
}
func (l *list) DecimalStyle() *list {
	l.style = decimal
	return l
}
func (l *list) HyphenStyle() *list {
	l.style = hyphen
	return l
}
func (l *list) DefaultStyle() *list {
	l.style = ""
	return l
}
