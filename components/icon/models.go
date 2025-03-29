package icon

import (
	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
)

type Icon interface {
	SetStroke(width int) Icon
	SetHeight(height int) Icon
	SetWidth(width int) Icon
	SetClsCustom(clsCustom string) Icon
	Component() templ.Component
	SetAttributes(attributes templ.Attributes) *component.C 
	GetAttributes() templ.Attributes
	GetClasses() []string
	AddClasses(styles ... string) *component.C
}
type icon struct {
	component.C
	name        string
	strokeWidth int
	height      int
	width       int
	clsCustom   string
}

// Create a new icon with default values, a lucid name is required, see https://lucide.dev/
func NewIcon(name string) Icon {
	i := &icon{name: name, strokeWidth: 2, height: 16, width: 16, clsCustom: " "}
	return i
}

func (i *icon) SetStroke(width int) Icon {
	i.strokeWidth = width
	return i
}

func (i *icon) SetHeight(height int) Icon {
	i.height = height
	return i
}

func (i *icon) SetWidth(width int) Icon {
	i.width = width
	return i
}

func (i *icon) SetClsCustom(clsCustom string) Icon {
	i.clsCustom = clsCustom
	return i
}
