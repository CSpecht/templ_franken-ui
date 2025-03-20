package typography

import (
	"github.com/a-h/templ"
)

type heading struct {
	content    string
	size       HeadingSize
	attributes templ.Attributes
	alignment  Alignment
	style      HeadingStyle
}

func NewHeading(content string) *heading {
	return &heading{
		content: content,
		size:    H1,
	}
}
func (h *heading) SetSize(size HeadingSize) *heading {
	h.size = size
	return h
}
func (h *heading) SetAttributes(attributes templ.Attributes) *heading {
	h.attributes = attributes
	return h
}
func (h *heading) SetAlignment(alignment Alignment) *heading {
	h.alignment = alignment
	return h
}
func (h *heading) SetHeadingStyle(style HeadingStyle) *heading {
	h.style = style
	return h
}
func (h *heading) Component() templ.Component {

	return h.create()
}

type hero struct {
	content    string
	size       HeroSize
	attributes templ.Attributes
}

func NewHero(content string) *hero {
	return &hero{
		content: content,
		size:    MD,
	}
}
func (he *hero) SetSize(size HeroSize) *hero {
	he.size = size
	return he
}
func (he *hero) SetAttributes(attributes templ.Attributes) *hero {
	he.attributes = attributes
	return he
}
func (he *hero) Component() templ.Component {

	return he.create()
}

type paragraph struct {
	content    string
	attributes templ.Attributes
}

func NewParagraph(content string) *paragraph {
	return &paragraph{
		content: content,
	}
}
func (p *paragraph) SetAttributes(attributes templ.Attributes) *paragraph {
	p.attributes = attributes
	return p
}
func (p *paragraph) Component() templ.Component {
	return p.create()
}

type blockquote struct {
	content    string
	attributes templ.Attributes
}

func NewBlockquote(content string) *blockquote {
	return &blockquote{
		content: content,
	}
}
func (b *blockquote) SetAttributes(attributes templ.Attributes) *blockquote {
	b.attributes = attributes
	return b
}
func (b *blockquote) Component() templ.Component {
	return b.create()
}

type inlineCode struct {
	content    string
	attributes templ.Attributes
}

func NewInlineCode(content string) *inlineCode {
	return &inlineCode{
		content: content,
	}
}
func (ic *inlineCode) SetAttributes(attributes templ.Attributes) *inlineCode {
	ic.attributes = attributes
	return ic
}
func (ic *inlineCode) Component() templ.Component {
	return ic.create()
}

type text struct {
	content    string
	attributes templ.Attributes
	alignment  Alignment
	size       TextSize
	textStyle  TextStyle
}

func NewText(content string) *text {
	return &text{
		content: content,
		size:    Default,
	}
}
func (t *text) SetAlignment(alignment Alignment) *text {
	t.alignment = alignment
	return t
}
func (t *text) SetSize(size TextSize) *text {
	t.size = size
	return t
}
func (t *text) SetTextStyle(style TextStyle) *text {
	t.textStyle = style
	return t
}
func (t *text) SetAttributes(attributes templ.Attributes) *text {
	t.attributes = attributes
	return t
}
func (t *text) Component() templ.Component {
	return t.create()
}
