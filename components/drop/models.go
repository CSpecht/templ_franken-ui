package drop

import (
	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
)

type triggerComponent interface {
	Component() templ.Component
	AddInnerComponent(comp templ.Component)
}
type drop struct {
	component.C
	triggerComponent triggerComponent
	targetComponent  templ.Component
	mode             string
	customParameters map[string]string

	
}

func NewDrop(triggerComponent triggerComponent, targetComponent templ.Component) *drop {
	return &drop{
		triggerComponent: triggerComponent,
		targetComponent:  targetComponent,
	}
}
func (drop *drop) SetCustomParameters(k,v string) *drop {
	if drop.customParameters == nil {
		drop.customParameters = make(map[string]string)
	}
	drop.customParameters[k] = v
	return drop
}
func (drop *drop) getCustomParameters() string {
	
	customParams := ""
	for k, v := range drop.customParameters {
		customParams += k + ":" + v + ";"
	}
	if drop.mode != "" {
		customParams += drop.mode + ";"
	}
	return customParams
}
func (drop *drop) ModeClickHover() *drop {
	drop.mode = "mode: click,hover"
	return drop
}
func (drop *drop) ModeClick() *drop {
	drop.mode = "mode: click"
	return drop
}
func (drop *drop) ModeHover() *drop {
	drop.mode = "mode: hover"
	return drop
}

func (drop *drop) AddParentIcon() *drop {
	parentIcon := drop.createParentIcon()
	drop.triggerComponent.AddInnerComponent(parentIcon)
	return drop
}

