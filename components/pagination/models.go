package pagination

import (
	"context"
	"io"

	"github.com/cspecht/templ_franken-ui/components/component"
)

type pagination struct {
	component.C
	style                        PaginationStyle
	size                         PaginationSize
	alignment                    PaginationAlignment
	elements                     []paginationElement
	previousLabel                string
	previousLabelAttributeStyles component.AttributesStyles
	nextLabel                    string
	nextLabelAttributeStyles     component.AttributesStyles
	skipPreviousNext             bool
}

func NewPagination() *pagination {
	return &pagination{
		style:            DefaultStyle,
		size:             Md,
		alignment:        DefaultAlignment,
		elements:         []paginationElement{},
		skipPreviousNext: false,
	}
}

func (p *pagination) WithPreviousLabel(label string, attributeStyles ...component.AttributesStyles) *pagination {
	p.previousLabel = label
	if len(attributeStyles) > 0 {
		p.previousLabelAttributeStyles = attributeStyles[0]
	}
	return p
}
func (p *pagination) WithNextLabel(label string, attributeStyles ...component.AttributesStyles) *pagination {
	p.nextLabel = label
	if len(attributeStyles) > 0 {
		p.nextLabelAttributeStyles = attributeStyles[0]
	}
	return p
}
func (p *pagination) SkipPreviousNext() *pagination {
	p.skipPreviousNext = true
	return p
}

func (p *pagination) WithStyle(style PaginationStyle) *pagination {
	p.style = style
	return p
}

func (p *pagination) WithSize(size PaginationSize) *pagination {
	p.size = size
	return p
}

func (p *pagination) WithAlignment(alignment PaginationAlignment) *pagination {
	p.alignment = alignment
	return p
}

type paginationElement struct {
	component.C
	active   bool
	disabled bool
	content  string
}

func (p *pagination) AddPage(content string, active bool, disabled bool, attributeStyles ...component.AttributesStyles) *pagination {
	element := paginationElement{
		active:   active,
		disabled: disabled,
		content:  content,
	}
	if active {
		element.SetClasses(string(Active))
	}
	if disabled {
		element.SetClasses(string(Disabled))
	}
	if len(attributeStyles) > 0 {
		element.SetAttributes(attributeStyles[0].Attributes)
		element.SetClasses(attributeStyles[0].Styles...)
	}
	p.elements = append(p.elements, element)
	return p
}
func (p *pagination) AddPlacHolder() *pagination {
	element := paginationElement{
		active:   false,
		disabled: true,
		content:  "...",
	}
	element.SetClasses(string(Disabled))
	p.elements = append(p.elements, element)
	return p
}

func (p *pagination) Render(ctx context.Context, w io.Writer) error {

	return p.component().Render(ctx, w)
}
