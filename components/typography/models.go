package typography

import (
	"context"
	"io"

	"github.com/cspecht/templ_franken-ui/components/component"
)

type heading struct {
	component.C
	content   string
	size      string
	alignment string
	style     string
}

func NewHeading(content string) *heading {
	return &heading{
		content: content,
		size:    h1,
	}
}
func (h *heading) Render(ctx context.Context, w io.Writer) error {

	return h.component().Render(ctx, w)
}

func (h *heading) SizeH1() *heading {
	h.size = h1
	return h
}
func (h *heading) SizeH2() *heading {
	h.size = h2
	return h
}
func (h *heading) SizeH3() *heading {
	h.size = h3
	return h
}
func (h *heading) SizeH4() *heading {
	h.size = h4
	return h
}
func (h *heading) AlignLeft() *heading {
	h.alignment = left
	return h
}
func (h *heading) AlignCenter() *heading {
	h.alignment = center
	return h
}
func (h *heading) AlignRight() *heading {
	h.alignment = right
	return h
}
func (h *heading) DividerStyle() *heading {
	h.style = divider
	return h
}
func (h *heading) LineStyle() *heading {
	h.style = line
	return h
}
func (h *heading) BulletStyle() *heading {
	h.style = bullet
	return h
}

type hero struct {
	component.C
	content string
	size    string
}

func NewHero(content string) *hero {
	return &hero{
		content: content,
		size:    md,
	}
}
func (h *hero) Render(ctx context.Context, w io.Writer) error {

	return h.component().Render(ctx, w)
}

func (he *hero) SizeSM() *hero {
	he.size = sm
	return he
}
func (he *hero) SizeMD() *hero {
	he.size = md
	return he
}
func (he *hero) SizeLG() *hero {
	he.size = lg
	return he
}
func (he *hero) SizeXL() *hero {
	he.size = xl
	return he
}
func (he *hero) Size2XL() *hero {
	he.size = twoXL
	return he
}
func (he *hero) Size3XL() *hero {
	he.size = threeXL
	return he
}

type paragraph struct {
	component.C
	content string
}

func NewParagraph(content string) *paragraph {
	return &paragraph{
		content: content,
	}
}
func (p *paragraph) Render(ctx context.Context, w io.Writer) error {

	return p.component().Render(ctx, w)
}

type blockquote struct {
	component.C
	content string
}

func NewBlockquote(content string) *blockquote {
	return &blockquote{
		content: content,
	}
}
func (bq *blockquote) Render(ctx context.Context, w io.Writer) error {

	return bq.component().Render(ctx, w)
}


type inlineCode struct {
	component.C
	content string
}

func NewInlineCode(content string) *inlineCode {
	return &inlineCode{
		content: content,
	}
}
func (ic *inlineCode) Render(ctx context.Context, w io.Writer) error {

	return ic.component().Render(ctx, w)
}


type text struct {
	component.C
	content   string
	alignment string
	size      string
	textStyle string
}

func NewText(content string) *text {
	return &text{
		content: content,
		size:    defaultText,
	}
}
func (t *text) Render(ctx context.Context, w io.Writer) error {

	return t.component().Render(ctx, w)
}

func (t *text) AlignLeft() *text {
	t.alignment = left
	return t
}
func (t *text) AlignCenter() *text {
	t.alignment = center
	return t
}
func (t *text) AlignRight() *text {
	t.alignment = right
	return t
}
func (t *text) SizeSmall() *text {
	t.size = small
	return t
}
func (t *text) SizeMedium() *text {
	t.size = medium
	return t
}
func (t *text) SizeLarge() *text {
	t.size = large
	return t
}
func (t *text) SizeDefault() *text {
	t.size = defaultText
	return t
}

func (t *text) Lead() *text {
	t.textStyle = lead
	return t
}
func (t *text) Meta() *text {
	t.textStyle = meta
	return t
}
func (t *text) Truncate() *text {
	t.textStyle = truncate
	return t
}
func (t *text) Breaks() *text {
	t.textStyle = breaks
	return t
}
