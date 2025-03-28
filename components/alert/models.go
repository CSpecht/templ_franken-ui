package alert
import (
	
	"github.com/cspecht/templ_franken-ui/components/component"
)

type alert struct{
	component.C
	content string
	title string
	closeable bool
	destructive bool
}

func NewAlert(content string) *alert {
	return &alert{
		content: content,
		closeable: false,
		destructive: false,
		title: "",
	}
}
func (a *alert) SetTitle(title string) *alert {
	a.title = title
	return a
}
func (a *alert) AsDestructive() *alert {
	a.destructive = true
	return a
}
func (a *alert) AsCloseable() *alert {
	a.closeable = true
	return a
}
func (a *alert) SetContent(content string) *alert {
	a.content = content
	return a
}