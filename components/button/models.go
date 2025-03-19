package button

import (
	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/icon"
)

type button struct {

	name          string
	classes       []string
	attributes    templ.Attributes
	buttonVariant buttonVariant
	buttonSize    buttonSize
	withModifiers withModifiers
	icon          *icon.Icon
}

func NewPrimaryButton(name string, classes []string, attr templ.Attributes, btnVariant buttonVariant, withModifiers withModifiers, icon *icon.Icon) *button {
	b := button{name: name, classes: classes, attributes: attr, buttonVariant: btnVariant, buttonSize: Sm, withModifiers: withModifiers, icon: icon}
	b.classes = append(b.classes, string(Primary))
	return &b
}

func NewSecondaryButton(name string, classes []string, attr templ.Attributes, btnVariant buttonVariant, buttonSize buttonSize, withModifiers withModifiers, icon *icon.Icon) *button {
	b := button{name: name, classes: classes, attributes: attr, buttonVariant: btnVariant, buttonSize: buttonSize, withModifiers: withModifiers, icon: icon}
	b.classes = append(b.classes, string(Secondary))
	return &b
}

func NewDefaultButton(name string, classes []string, attr templ.Attributes, btnVariant buttonVariant, buttonSize buttonSize, withModifiers withModifiers, icon *icon.Icon) *button {
	b := button{name: name, classes: classes, attributes: attr, buttonVariant: btnVariant, buttonSize: buttonSize, withModifiers: withModifiers, icon: icon}
	b.classes = append(b.classes, string(Default))
	return &b
}

func NewGhostButton(name string, classes []string, attr templ.Attributes, btnVariant buttonVariant, buttonSize buttonSize, withModifiers withModifiers, icon *icon.Icon) *button {
	b := button{name: name, classes: classes, attributes: attr, buttonVariant: btnVariant, buttonSize: buttonSize, withModifiers: withModifiers, icon: icon}
	b.classes = append(b.classes, string(Ghost))
	return &b
}

func NewDestructiveButton(name string, classes []string, attr templ.Attributes, btnVariant buttonVariant, buttonSize buttonSize, withModifiers withModifiers, icon *icon.Icon) *button {
	b := button{name: name, classes: classes, attributes: attr, buttonVariant: btnVariant, buttonSize: buttonSize, withModifiers: withModifiers, icon: icon}
	b.classes = append(b.classes, string(Destructive))
	return &b
}

func NewTextButton(name string, classes []string, attr templ.Attributes, btnVariant buttonVariant, buttonSize buttonSize, withModifiers withModifiers, icon *icon.Icon) *button {
	b := button{name: name, classes: classes, attributes: attr, buttonVariant: btnVariant, buttonSize: buttonSize, withModifiers: withModifiers, icon: icon}
	b.classes = append(b.classes, string(Text))
	return &b
}

func NewLinkButton(name string, classes []string, attr templ.Attributes, btnVariant buttonVariant, buttonSize buttonSize, withModifiers withModifiers, icon *icon.Icon) *button {
	b := button{name: name, classes: classes, attributes: attr, buttonVariant: btnVariant, buttonSize: buttonSize, withModifiers: withModifiers, icon: icon}
	b.classes = append(b.classes, string(Link))
	return &b
}

func (b *button) AddClass(class string) *button {
	b.classes = append(b.classes, class)
	return b
}
func (b *button) SetSize(size buttonSize) *button {
	b.buttonSize = size
	return b
}
func (b *button) Component() templ.Component {
	return b.create()
}
