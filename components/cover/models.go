package cover

import (
	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
	"github.com/cspecht/templ_franken-ui/components/video"
)

type imageCover struct {
	component.C
	src        templ.SafeURL
	responsive bool
	alt        string
}

func NewImageCover(src templ.SafeURL) *imageCover {
	return &imageCover{
		src: src,
	}
}
func (i *imageCover) Responsive() *imageCover {
	i.responsive = true
	return i
}
func (i *imageCover) SetAlt(alt string) *imageCover {
	i.alt = alt
	return i
}

type videoCover struct {
	component.C
	video      video.Video
	responsive bool
}

func NewVideoCover(v video.Video) *videoCover {
	vc := &videoCover{video: v}
	attr := vc.C.GetAttributes()
	if attr == nil {
		attr = make(templ.Attributes)
	}
	attr["data-uk-cover"] = true
	vc.C.SetAttributes(attr)
	return vc

}

func (v *videoCover) Responsive() *videoCover {
	v.responsive = true
	return v
}
func (v *videoCover) SetVideo(vd video.Video) *videoCover {
	v.video = vd
	return v
}

type iframeCover struct {
	component.C
	src        templ.SafeURL
	responsive bool
}

func NewIframeCover(src templ.SafeURL) *iframeCover {
	return &iframeCover{
		src: src,
	}
}
func (i *iframeCover) Responsive() *iframeCover {
	i.responsive = true
	return i
}
func (i *iframeCover) SetSrc(src templ.SafeURL) *iframeCover {
	i.src = src
	return i
}
