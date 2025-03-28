package placeholder

import (
	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
)

type placeholder struct {
	component.C
	alignment string
	content   []templ.Component
}

func NewPlaceholder(content ...templ.Component) *placeholder {
	return &placeholder{
		content:   content,
		alignment: center,
	}
}

func (p *placeholder) AlignRight() *placeholder {
	p.alignment = right
	return p
}
func (p *placeholder) AlignLeft() *placeholder {
	p.alignment = left
	return p
}
func (p *placeholder) AlignCenter() *placeholder {
	p.alignment = center
	return p
}
func (p *placeholder) SetContent(content ...templ.Component) *placeholder {
	p.content = content
	return p
}


