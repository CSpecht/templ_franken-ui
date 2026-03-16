package cover

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
	"github.com/cspecht/templ_franken-ui/components/video"
)

type imageCover struct {
	component.C
	src        templ.SafeURL
	responsive bool
	alt        string
	width      int
	height     int
}

func NewImageCover(src templ.SafeURL) *imageCover {
	return &imageCover{
		src: src,
	}
}

// To add responsive behavior to your cover image, you need to create an invisible <canvas> element and assign width and height values to it, according to the aspect ratio you want the covered area to have. That way it will adapt the responsive behavior of the image.
func (i *imageCover) Responsive(width, height int) *imageCover {
	i.responsive = true
	i.width = width
	i.height = height
	return i
}
func (i *imageCover) SetAlt(alt string) *imageCover {
	i.alt = alt
	return i
}
func (i *imageCover) Render(ctx context.Context, w io.Writer) error {

	return i.component().Render(ctx, w)
}

type videoCover struct {
	component.C
	video      video.Video
	responsive bool
	width      int
	height     int
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

func (v *videoCover) Responsive(width, height int) *videoCover {
	v.responsive = true
	v.width = width
	v.height = height
	return v
}
func (v *videoCover) SetVideo(vd video.Video) *videoCover {
	v.video = vd
	return v
}
func (v *videoCover) Render(ctx context.Context, w io.Writer) error {

	return v.component().Render(ctx, w)
}

type iframeCover struct {
	component.C
	src        templ.SafeURL
	responsive bool
	width      int
	height     int
}

func NewIframeCover(src templ.SafeURL) *iframeCover {
	return &iframeCover{
		src: src,
	}
}
func (i *iframeCover) Responsive(width, height int) *iframeCover {
	i.responsive = true
	i.width = width
	i.height = height
	return i
}
func (i *iframeCover) SetSrc(src templ.SafeURL) *iframeCover {
	i.src = src
	return i
}
func (i *iframeCover) Render(ctx context.Context, w io.Writer) error {

	return i.component().Render(ctx, w)
}
