package image

import (
	"context"
	"encoding/json"
	"io"

	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
)

type image struct {
	component.C
	dataSrc string
	sources map[string]string
	sizes   string
	loading string
	margin  string
	target  string
	content templ.Component
}

func NewImage(src string) *image {
	return &image{dataSrc: src, sources: make(map[string]string), loading: "lazy", margin: "50%", target: "false"}
}

func (i *image) AddSources(sources map[string]string) *image {
	i.sources = sources
	return i
}
func (i *image) AddSizes(sizes string) *image {
	i.sizes = sizes
	return i
}
func (i *image) getParameters() string {
	parameters := make(map[string]string)
	if i.loading != "" {
		parameters["loading"] = i.loading
	}
	if i.margin != "" {
		parameters["margin"] = i.margin
	}
	if i.target != "" {
		parameters["target"] = i.target
	}
	jsonData, err := json.Marshal(parameters)
	if err != nil {
		return ""
	}
	var params string
	err = json.Unmarshal(jsonData, &params)
	if err != nil {
		return ""
	}
	return params
}

// todo error handling
func (i *image) getSources() string {

	jsonData, err := json.Marshal(i.sources)
	if err != nil {
		return ""
	}
	var srcset string
	err = json.Unmarshal(jsonData, &srcset)
	if err != nil {
		return ""
	}
	return srcset
}

func (i *image) SetLazyLoading() *image {
	i.loading = "lazy"
	return i
}

func (i *image) SetEagerLoading() *image {
	i.loading = "eager"
	return i
}

func (i *image) SetMargin(margin string) *image {
	i.margin = margin
	return i
}

func (i *image) SetTarget(target string) *image {
	i.target = target
	return i
}
func (i *image) SetContent(content templ.Component) *image {
	i.content = content
	return i
}
func (i *image) Render(ctx context.Context, w io.Writer) error {

	return i.component().Render(ctx, w)
}
