package divider

import "github.com/a-h/templ"

type divider struct {
	classes []string
	variant variant
}

func NewDivider() *divider {
	d := divider{classes: []string{}}

	return &d
}

func (d *divider) AsIcon() *divider {
	d.variant = Icon
	return d
}

func (d *divider) AsSmall() *divider {
	d.variant = Small
	return d
}

func (d *divider) AsVertical() *divider {
	d.variant = Vertical
	return d
}

func (d *divider) Component() templ.Component {
	d.classes = append(d.classes, string(d.variant))
	return d.create()
}
