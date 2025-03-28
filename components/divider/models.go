package divider

import 	"github.com/cspecht/templ_franken-ui/components/component"

type divider struct {
	component.C
	style string
	
}

func NewDivider() *divider {
	return  &divider{
		style: defaultVariant,
	}

}

func (d *divider) WithIcon() *divider {
	d.style = icon
	return d
}

func (d *divider) Small() *divider {
	d.style = small
	return d
}

func (d *divider) Vertical() *divider {
	d.style = vertical
	return d
}

