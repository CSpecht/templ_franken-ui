package button

import (
	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/icon"
	//"github.com/cspecht/templ_franken-ui/components/icon"
)

type button struct {
	name       string
	classes    []string
	attributes templ.Attributes
	variant    variant
	size       size
	width      width
	icon       *icon.Icon
	url        string
}

// classes []string, attr templ.Attributes, btnVariant buttonVariant, withModifiers withModifiers, icon *icon.Icon
// without size == default
func NewButton(name string) *button {
	b := button{name: name, classes: []string{}, attributes: make(templ.Attributes), variant: Default, width: W40, icon: nil}

	return &b
}

func (b *button) AddClass(class string) *button {
	b.classes = append(b.classes, class)
	return b
}

func (b *button) AddAttr(attr templ.Attributes) *button {
	for k, v := range b.attributes {
		print(k, v)
	}
	return b
}

func (b *button) SetSize(size size) *button {
	b.size = size
	return b
}
func (b *button) SetIcon(icon *icon.Icon) *button {
	b.icon = icon
	//doesnt work!
	//b.classes = append(b.classes, "uk-btn-icon")
	return b
}

func (b *button) SetWidth(width width) *button {
	b.width = width
	return b
}

func (b *button) AsPrimary() *button {
	b.variant = Primary
	return b
}

func (b *button) AsDefault() *button {
	b.variant = Default
	return b
}

func (b *button) AsGhost() *button {
	b.variant = Ghost
	return b
}
func (b *button) AsSecondary() *button {
	b.variant = Secondary
	return b
}
func (b *button) AsDestructive() *button {
	b.variant = Destructive
	return b
}

func (b *button) AsText() *button {
	b.variant = Text
	return b
}

func (b *button) AsLink(url string) *button {
	//b.url
	b.variant = Link
	return b
}

func (b *button) Component() templ.Component {
	b.classes = append(b.classes, string(b.variant))
	b.classes = append(b.classes, string(b.size))
	b.classes = append(b.classes, string(b.width))

	if len(b.url) > 0 {
		b.classes = append(b.classes, string(Link))
	}

	return b.create()
}
