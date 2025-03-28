package button

import (
	"github.com/cspecht/templ_franken-ui/components/component"
	"github.com/cspecht/templ_franken-ui/components/icon"
)

type button struct {
	component.C
	name    string
	variant string
	size    string
	width   string
	icon    icon.Icon
}

type buttonGroup struct {
	component.C
	items []*button
}

func NewButtonGroup() *buttonGroup {
	return &buttonGroup{}
}
func (bg *buttonGroup) AddButton(b *button) *buttonGroup {
	bg.items = append(bg.items, b)
	return bg
}

func NewButton(name string) *button {
	b := button{name: name,
		variant: defaultVariant, width: w40, icon: nil}

	return &b
}

func (b *button) SetSizeXs() *button {
	b.size = xs
	return b
}
func (b *button) SetSizeSm() *button {
	b.size = sm
	return b
}
func (b *button) SetSizeMd() *button {
	b.size = md
	return b
}
func (b *button) SetSizeLg() *button {
	b.size = lg
	return b
}
func (b *button) SetWidth40() *button {
	b.width = w40
	return b
}
func (b *button) SetWidth44() *button {
	b.width = w44
	return b
}
func (b *button) SetWidth48() *button {
	b.width = w48
	return b
}
func (b *button) SetWidth52() *button {
	b.width = w52
	return b
}
func (b *button) SetWidthFull() *button {
	b.width = wFull
	return b
}

func (b *button) SetIcon(icon icon.Icon) *button {
	b.icon = icon

	return b
}

func (b *button) AsPrimary() *button {
	b.variant = primary
	return b
}

func (b *button) AsDefault() *button {
	b.variant = defaultVariant
	return b
}

func (b *button) AsGhost() *button {
	b.variant = ghost
	return b
}
func (b *button) AsSecondary() *button {
	b.variant = secondary
	return b
}
func (b *button) AsDestructive() *button {
	b.variant = destructive
	return b
}

func (b *button) AsText() *button {
	b.variant = text
	return b
}

func (b *button) AsLink() *button {
	b.variant = link
	return b
}
