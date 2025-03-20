package placeholder

import "github.com/a-h/templ"


type placeholder struct {
	alignment string
	content []templ.Component
	attributes templ.Attributes
}


func NewPlaceholder(content ...templ.Component) *placeholder {
	return &placeholder{
		content: content,
		alignment: string(Center),
	}
}

func (p *placeholder) SetAlignment(alignment alignment) *placeholder {
	p.alignment = string(alignment)
	return p
}
func (p *placeholder) SetAttributes(attributes templ.Attributes) *placeholder {
	p.attributes = attributes
	return p
}

func (p *placeholder) Component() templ.Component {

	return p.create()
}
