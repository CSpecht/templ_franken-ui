package leader

import (
	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
)

type leader struct {
	component.C
	leftText string
	leftComponent templ.Component
	rightText string
	rightComponent templ.Component

	customParameters map[string]string
}
func NewLeader(leftText string, rightText string) *leader {
	l := &leader{
		leftText: leftText,
		rightText: rightText,
		
		customParameters: make(map[string]string),
		
	}
	l.customParameters["fill"] = "."
	return l
}
func (l *leader) SetLeftComponent(comp templ.Component) *leader {
	l.leftComponent = comp
	return l
}
func (l *leader) SetRightComponent(comp templ.Component) *leader {
	l.rightComponent = comp
	return l
}
func (l *leader) SetFillCharacter(char string) *leader {
	l.customParameters["fill"] = char
	return l
}
// see https://franken-ui.dev/docs/2.0/leader#responsive
func (l *leader) SetResponsive(responsive string) *leader {
	l.customParameters["media"] = responsive
	return l
}
func (l *leader) getCustomParameters() string {
	
	customParams := ""
	for k, v := range l.customParameters {
		customParams += k + ":" + v + ";"
	}

	return customParams
}