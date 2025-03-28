package badge

import "github.com/cspecht/templ_franken-ui/components/component"

type badge struct {
	component.C
	content string
	style   variant
}

func NewBadge(content string) *badge {
	return &badge{}

}

func (d *badge) AsPrimary() *badge {
	d.style = primary
	return d
}

func (d *badge) AsSecondary() *badge {
	d.style = secondary
	return d
}

func (d *badge) AsDestructive() *badge {
	d.style = destructive
	return d
}
