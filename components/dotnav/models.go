package dotnav

import "github.com/cspecht/templ_franken-ui/components/component"

type dotnav struct {
	component.C
	items []*dotnavItem
	vertical bool
}
type dotnavItem struct {
	component.C
	name string
	active bool
}
func NewDotnav() *dotnav {
	return &dotnav{}
}
func (d *dotnav) AddItem(name string) *dotnavItem {
	item := &dotnavItem{name: name}
	d.items = append(d.items, item)
	return item
}
func (di *dotnavItem) Active() *dotnavItem {
	di.active = true
	return di
}
func (di *dotnavItem) Inactive() *dotnavItem {
	di.active = false
	return di
}
func (d *dotnav) SetVertical() *dotnav {
	d.vertical = true
	return d
}
func (d *dotnav) SetHorizontal() *dotnav {
	d.vertical = false
	return d
}
func (d *dotnav) GetItems() []*dotnavItem {
	return d.items
}