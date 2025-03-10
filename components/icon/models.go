package icon

import "github.com/a-h/templ"

type Icon struct {
	name string
}

func NewIcon(name string) *Icon {
	i := Icon{name: name}
	return &i
}

func (i *Icon) Component() templ.Component {
	return i.create()

}
