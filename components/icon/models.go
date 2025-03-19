package icon

import "github.com/a-h/templ"

type Icon struct {
	name        string
	strokeWidth int
	height      int
	width       int
	clsCustom   string
}

func NewIcon(name string) *Icon {
	i := &Icon{name: name, strokeWidth: 2, height: 16, width: 16, clsCustom: " "}
	return i
}

func (i *Icon) SetStroke(width int) *Icon {
	i.strokeWidth = width
	return i
}

func (i *Icon) SetHeight(height int) *Icon {
	i.height = height
	return i
}

func (i *Icon) SetWidth(width int) *Icon {
	i.width = width
	return i
}

func (i *Icon) SetClsCustom(clsCustom string) *Icon {
	i.clsCustom = clsCustom
	return i
}

func (i *Icon) Component() templ.Component {
	return i.create()

}
