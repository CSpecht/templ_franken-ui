package video

import (
	"strings"

	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
)
type Video interface {
	
	EnableControls() *video
	EnablePlaysinline() *video
	EnableHidden() *video
	AllowFullscreen() *video
	SetSize(width, height int) *video
	EnableAutoplayInview() *video
	EnableAutoplayHover() *video
	EnableAutomute() *video
	getVideoAttributes() string
	SetAttributes(attributes templ.Attributes) *component.C 
	GetAttributes() templ.Attributes
	GetClasses() []string
	AddClasses(styles ... string) *component.C
	Component() templ.Component
}

type video struct {
	component.C
	src templ.SafeURL
	controls bool
	playsinline bool
	hidden bool
	autoplay string
	automute string
	width int
	height int
	allowFullscreen bool
}

func NewVideo(src templ.SafeURL) *video {
	return &video{
		src: src,
	}
}
func (v *video)EnableControls() *video {
	v.controls = true
	attrs := v.C.GetAttributes()
	if attrs == nil {
		attrs = make(templ.Attributes)
	}
	attrs["controls"] = true
	v.C.SetAttributes(attrs)
	return v
}
func (v *video)EnablePlaysinline() *video {
	v.playsinline = true
	attrs := v.C.GetAttributes()
	if attrs == nil {
		attrs = make(templ.Attributes)
	}
	attrs["playsinline"] = true
	v.C.SetAttributes(attrs)
	return v
}
func (v *video)EnableHidden() *video {
	v.hidden = true
	attrs := v.C.GetAttributes()
	if attrs == nil {
		attrs = make(templ.Attributes)
	}
	attrs["hidden"] = true
	v.C.SetAttributes(attrs)
	return v
}
func (v *video) AllowFullscreen() *video {
	v.allowFullscreen = true
	attrs := v.C.GetAttributes()
	if attrs == nil {
		attrs = make(templ.Attributes)
	}
	attrs["allowfullscreen"] = true
	v.C.SetAttributes(attrs)
	return v
}
func (v *video)SetSize(width, height int) *video {
	v.width = width
	v.height = height
	return v
}
func (v *video)EnableAutoplayInview() *video {
	v.autoplay = "autoplay: inview"
	return v
}
func (v *video)EnableAutoplayHover() *video {
	v.autoplay = "autoplay: hover"
	return v
}
func (v *video)EnableAutomute() *video {
	v.automute = "automute: true"
	return v
}
func(v *video)getVideoAttributes()string{
	l := []string{}
	if len(v.autoplay) > 0 {
		l = append(l, v.autoplay)
	}
	if len(v.automute) > 0 {
		l = append(l, v.automute)
	}
	return strings.Join(l, "; ")
}