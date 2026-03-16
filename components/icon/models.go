package icon

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
)

type Icon interface {
	SetStroke(width int) Icon
	SetHeight(height int) Icon
	SetWidth(width int) Icon
	SetClsCustom(clsCustom string) Icon
	SetAttributes(attributes templ.Attributes) *component.C
	GetAttributes() templ.Attributes
	GetClasses() []string
	SetClasses(styles ...string) *component.C
	Render(ctx context.Context, w io.Writer) error
}
type icon struct {
	component.C
	name                 string
	strokeWidth          int
	height               int
	width                int
	clsCustom            string
	forcePreventReRender bool
}

// Create a new icon with default values, a lucid name is required, see https://lucide.dev/
func NewIcon(name string) *icon {
	i := &icon{name: name, strokeWidth: 2, height: 16, width: 16, forcePreventReRender: false, clsCustom: " "}
	return i
}
func (i *icon) Render(ctx context.Context, w io.Writer) error {

	return i.component().Render(ctx, w)
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
func (i *icon) ForcePreventReRender(prevent bool) Icon {
	i.forcePreventReRender = prevent
	return i
}

// Allows you to add custom CSS classes, which will be attached to the <svg> tag.
func (i *icon) SetClsCustom(clsCustom string) Icon {
	i.clsCustom = clsCustom
	return i
}
