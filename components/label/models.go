package label

import (
	"github.com/cspecht/templ_franken-ui/components/component"
)

type label struct {
	component.C
	content string
	style   string
}

func NewLabel(content string) *label {
	return &label{
		content: content,
		style:   "",
	}
}
func (l *label) AsDestructive() *label {
	l.style = destructive
	return l
}
func (l *label) AsDefault() *label {
	l.style = ""
	return l
}
func (l *label) AsPrimary() *label {
	l.style = primary
	return l
}
func (l *label) AsSecondary() *label {
	l.style = secondary
	return l
}
