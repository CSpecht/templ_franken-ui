package label

import (
	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
)

type label struct {
	component.C
	content string
	style   LabelStyle
}

func NewLabel(content string) *label {
	return &label{
		content: content,
		style:   Default,
	}
}
func (l *label) AsDestructive() *label {
	l.style = Destructive
	return l
}
func (l *label) AsDefault() *label {
	l.style = Default
	return l
}
func (l *label) AsPrimary() *label {
	l.style = Primary
	return l
}
func (l *label) AsSecondary() *label {
	l.style = Secondary
	return l
}
func (l *label) SetStyle(style LabelStyle) *label {
	l.style = style
	return l
}
func (l *label) Component() templ.Component {
	return l.create()
}
