package table

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
)

type table struct {
	component.C
	caption      string
	headElements []*cellElement
	rows         []*rowElement
	footElements []*cellElement
	responsive   TableModifier
}

func NewTable(attrStyles ...component.AttributesStyles) *table {
	t := table{}
	if len(attrStyles) > 0 {
		t.SetAttributes(attrStyles[0].Attributes)
		t.SetClasses(attrStyles[0].Styles...)
	} else{
		t.SetAttributes(make(templ.Attributes))
	}
	return &t
}
func (t *table) WithCaption(caption string) *table {
	t.caption = caption
	return t
}
func (t *table) IsResponsiveOverflow() *table {
	t.responsive = ResponsiveOverflow
	
	return t
}
func (t *table) IsResponsiveStack() *table {
	t.responsive = ResponsiveStack
	t.SetClasses(ResponsiveStack.String())
	return t
}
func (t *table) WithDivider()*table{
	t.SetClasses(Divider.String())
	return t
}
func (t *table) IsStriped()*table{
	t.SetClasses(Striped.String())
	return t
}
func (t *table) WithHover()*table{
	t.SetClasses(Hover.String())
	return t
}
func (t *table) SetSize(size TableSize)*table{
	t.SetClasses(size.String())
	return t
}
func (t *table)AddTableModifier(tm TableModifier)*table{
	t.SetClasses(tm.String())
	return t
}
// To remove the outer padding of the first and last columns so that they are flush with the table
func (t *table)Justify()*table{
	t.SetClasses(Justify.String())
	return t
}
func (t *table) AddHeader(headers ...*cellElement) *table {
	t.headElements = append(t.headElements, headers...)
	return t
}
func (t *table) AddFooter(footers ...*cellElement) *table {
	t.footElements = append(t.footElements, footers...)
	return t
}

type rowElement struct {
	component.C
	cells []*cellElement
}

func (t *table) NewRowElement(attrStyles ...component.AttributesStyles) *rowElement {
	row := &rowElement{}
	if len(attrStyles) > 0 {
		row.SetAttributes(attrStyles[0].Attributes)
		row.SetClasses(attrStyles[0].Styles...)
	} else{
		row.SetAttributes(make(templ.Attributes))
	}
	t.rows = append(t.rows, row)
	return row

}
func (r *rowElement) AddCells(cells ...*cellElement) *rowElement {
	r.cells = append(r.cells, cells...)
	return r
}
func (r *rowElement) AddTableModifier(tm TableModifier)*rowElement{
	r.SetClasses(tm.String())
	return r
}

type cellElement struct {
	component.C
	content templ.Component
	text    string
}

func (t *table) NewCellElement(comp templ.Component, attrStyles ...component.AttributesStyles) *cellElement {
	cell := cellElement{}
	if len(attrStyles) > 0 {
		cell.SetAttributes(attrStyles[0].Attributes)
		cell.SetClasses(attrStyles[0].Styles...)
	} else{
		cell.SetAttributes(make(templ.Attributes))
	}
	return &cell
}
func (t *table) NewCellElementText(text string, attrStyles ...component.AttributesStyles) *cellElement {
	cell := cellElement{}
	cell.text = text
	if len(attrStyles) > 0 {
		cell.SetAttributes(attrStyles[0].Attributes)
		cell.SetClasses(attrStyles[0].Styles...)
	}else{
		cell.SetAttributes(make(templ.Attributes))
	}
	return &cell
}
func (c *cellElement)AddTableModifier(tm TableModifier)*cellElement{
	c.SetClasses(tm.String())
	return c
}
func (t *table) getHeaderTitle(idx int, attr templ.Attributes) string {
	// first we should check if they key exist
	if attr !=nil{
	value, ok := attr["data-label"]
		if ok {
			
			// can parse?
			sValue, ok := value.(string)
			if ok{
				return sValue
			}
		} 
	}
	
	if idx >= 0 && idx < len(t.headElements) {
		return t.headElements[idx].text

	} else {
		return ""
	}
}

func (t *table) Render(ctx context.Context, w io.Writer) error {
	if t.responsive == ResponsiveOverflow {
		return t.divWrapper(t.component()).Render(ctx, w)
	}
	return t.component().Render(ctx, w)
}
