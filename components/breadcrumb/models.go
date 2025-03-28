package breadcrumb

import (
	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
)
type breadcrumb struct {
	component.C
	elements []templ.Component
}

func NewBreadcrumb() *breadcrumb {
	return &breadcrumb{
		
	}
}

func (b *breadcrumb) AddCrumb(name, link string,attributes ...templ.Attributes ) *breadcrumb {
	b.elements = append(b.elements, b.addCrumb(name, link, attributes...))
	return b
}
func (b *breadcrumb) AddDisabledCrumb(name string) *breadcrumb {
	b.elements = append(b.elements, b.addDisabledCrumb(name))
	return b
}
func (b *breadcrumb) AddCurrentCrumb (name string) *breadcrumb {
	b.elements = append(b.elements, b.addCurrentCrumb(name))
	return b
}