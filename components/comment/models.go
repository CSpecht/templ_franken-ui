package comment

import (
	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
)

type comment struct {
	component.C
	primary        bool
	avatarImageUrl templ.SafeURL

	avatarAlt   string
	title       string
	titleLink   templ.SafeURL
	meta        string
	metaLink    templ.SafeURL
	text        string
	body        templ.Component
	commentlist *commentList
}

func NewComment(title, meta string, text ...string) *comment {
	c := &comment{
		title: title,
		meta:  meta,
	}
	if len(text) > 0 {
		c.text = text[0]
	}
	return c

}
func (c *comment) AsPrimary() *comment {
	c.primary = true
	return c
}
func (c *comment) SetAvatarImageUrl(url templ.SafeURL) *comment {
	c.avatarImageUrl = url
	return c
}
func (c *comment) SetAvatarAlt(alt string) *comment {
	c.avatarAlt = alt
	return c
}
func (c *comment) SetTitle(title string) *comment {
	c.title = title
	return c
}
func (c *comment) SetTitleLink(link templ.SafeURL) *comment {
	c.titleLink = link
	return c
}
func (c *comment) SetMeta(meta string) *comment {
	c.meta = meta
	return c
}
func (c *comment) SetMetaLink(link templ.SafeURL) *comment {
	c.metaLink = link
	return c
}
func (c *comment) SetText(text string) *comment {
	c.text = text
	return c
}
func (c *comment) SetBody(body templ.Component) *comment {
	c.body = body
	return c
}

func (c *comment) AddComment(c2 *comment) *comment {
	if c.commentlist == nil {
		c.commentlist = NewCommentList()
	}
	c.commentlist.AddComment(c2)
	return c
}

type commentList struct {
	component.C
	comments []*comment
}

func NewCommentList() *commentList {
	return &commentList{}
}
func (cl *commentList) AddComment(c *comment) *commentList {
	cl.comments = append(cl.comments, c)
	return cl
}
