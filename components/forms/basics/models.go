package basics

import (
	"context"
	"io"

	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/component"
	"github.com/cspecht/templ_franken-ui/components/icon"
)

type form struct {
	component.C
	comps            []templ.Component
	layout           FormLayout
	action           templ.Component
	actionAttributes templ.Attributes
	actionStyles     []string
}

func NewForm(attrStyles ...component.AttributesStyles) *form {
	f := form{}
	if len(attrStyles) > 0 {
		f.SetClasses(attrStyles[0].Styles...)
		if attrStyles[0].Attributes != nil {
			f.SetAttributes(attrStyles[0].Attributes)
		} else {
			f.SetAttributes(make(templ.Attributes))
		}
	}
	return &f
}

// Use the uk-form-controls-text class to better align checkboxes and radio buttons when using them with text in a horizontal layout. Use NewFormControl...IsHorizontal() helper function
func (f *form) WithFormlayout(lay FormLayout) *form {
	f.layout = lay
	return f

}
func (f *form) AddComponents(comps ...templ.Component) *form {
	f.comps = append(f.comps, comps...)
	return f
}
func (f *form) AddActions(comp templ.Component, attrStyles ...component.AttributesStyles) *form {
	f.action = comp
	if len(attrStyles) > 0 {
		f.actionStyles = attrStyles[0].Styles
		if attrStyles[0].Attributes != nil {
			f.actionAttributes = attrStyles[0].Attributes
		} else {
			f.actionAttributes = make(templ.Attributes)
		}
	}
	return f
}

func (f *form) Render(ctx context.Context, w io.Writer) error {
	return f.component().Render(ctx, w)
}

type fieldSet struct {
	component.C
	legend string
	comps  []templ.Component
	layout FormLayout
}

func NewFieldSet(attrStyles ...component.AttributesStyles) *fieldSet {
	fs := fieldSet{}
	if len(attrStyles) > 0 {
		fs.SetClasses(attrStyles[0].Styles...)
		if attrStyles[0].Attributes != nil {
			fs.SetAttributes(attrStyles[0].Attributes)
		} else {
			fs.SetAttributes(make(templ.Attributes))
		}
	}
	return &fs
}
func (fs *fieldSet) WithLegend(leg string) *fieldSet {
	fs.legend = leg
	return fs

}

// Use the uk-form-controls-text class to better align checkboxes and radio buttons when using them with text in a horizontal layout. Use NewFormControl...IsHorizontal() helper function
func (fs *fieldSet) WithFormlayout(lay FormLayout) *fieldSet {
	fs.layout = lay
	return fs

}

func (fs *fieldSet) AddComponents(comps ...templ.Component) *fieldSet {
	fs.comps = append(fs.comps, comps...)
	return fs
}

func (fs *fieldSet) Render(ctx context.Context, w io.Writer) error {
	return fs.component().Render(ctx, w)
}

type formControl struct {
	component.C
	comp       templ.Component
	horizontal bool
}

func NewFormControls(formElement templ.Component, attrStyles ...component.AttributesStyles) *formControl {
	fc := formControl{comp: formElement}
	if len(attrStyles) > 0 {
		fc.SetClasses(attrStyles[0].Styles...)
		if attrStyles[0].Attributes != nil {
			fc.SetAttributes(attrStyles[0].Attributes)
		} else {
			fc.SetAttributes(make(templ.Attributes))
		}
	}
	return &fc
}
func (fc *formControl) IsHorizontal() *formControl {
	fc.horizontal = true
	return fc
}

func (fc *formControl) Render(ctx context.Context, w io.Writer) error {
	return fc.component().Render(ctx, w)
}

type label struct {
	component.C
	name     string
	id       string
	labelFor string
	required bool
	hasError bool
}

func NewLabel(name, labelFor string) *label {
	l := label{
		name:     name,
		labelFor: labelFor,
	}
	return &l
}
func (l *label) WithId(id string) *label {
	l.id = id
	return l
}
func (l *label) IsRequired() *label {
	l.required = true
	return l
}
func (l *label) HasError() *label {
	l.hasError = true
	return l
}
func (l *label) Render(ctx context.Context, w io.Writer) error {
	return l.component().Render(ctx, w)
}

type inputText struct {
	component.C
	t               InputType
	id              string
	name            string
	value           string
	required        bool
	minLength       int
	maxLength       int
	size            int
	pattern         string
	placeholder     string
	readonly        bool
	spellcheck      bool
	autocomplete    bool
	disabled        bool
	hasError        bool
	ariaDescribedby string
	ariaLabel       string
	sizeModifier    SizeModifier
}

func NewInputText(name, id, value string, attrStyles ...component.AttributesStyles) *inputText {
	it := inputText{
		t:            Text,
		name:         name,
		id:           id,
		minLength:    -1,
		maxLength:    -1,
		size:         -1,
		spellcheck:   true,
		autocomplete: true,
	}
	if len(attrStyles) > 0 {
		it.SetClasses(attrStyles[0].Styles...)
		if attrStyles[0].Attributes != nil {
			it.SetAttributes(attrStyles[0].Attributes)
		} else {
			it.SetAttributes(make(templ.Attributes))
		}
	}

	return &it
}

func (it *inputText) WithSize(size SizeModifier) *inputText {
	it.sizeModifier = size
	return it

}
func (it *inputText) IsRequired() *inputText {
	it.required = true
	return it
}
func (it *inputText) WithLenght(min, max int) *inputText {
	it.minLength = min
	it.maxLength = max
	return it
}
func (it *inputText) WithInputSize(s int) *inputText {
	it.size = s
	return it
}
func (it *inputText) WithPattern(pat string) *inputText {
	it.pattern = pat
	return it
}
func (it *inputText) WithPlacerholder(ph string) *inputText {
	it.placeholder = ph
	return it
}
func (it *inputText) IsReadonly() *inputText {
	it.readonly = true
	return it
}
func (it *inputText) DisableSpellcheck() *inputText {
	it.spellcheck = false
	return it
}
func boolOnOff(act bool) string {
	if act {
		return "on"
	} else {
		return "off"
	}
}
func boolTrueFalse(bl bool) string {
	if bl {
		return "true"
	} else {
		return "false"
	}
}
func (it *inputText) DisableAutocomplete() *inputText {
	it.autocomplete = false
	return it
}
func (it *inputText) Render(ctx context.Context, w io.Writer) error {
	return it.component().Render(ctx, w)
}

type formIcon struct {
	component.C
	formIcon    icon.Icon
	isFlipped   bool
	isClickable bool
}

func NewFormIcon(ic icon.Icon, attrStyles ...component.AttributesStyles) *formIcon {
	fi := formIcon{formIcon: ic}
	if len(attrStyles) > 0 {
		fi.SetClasses(attrStyles[0].Styles...)
		if attrStyles[0].Attributes != nil {
			fi.SetAttributes(attrStyles[0].Attributes)
		} else {
			fi.SetAttributes(make(templ.Attributes))
		}
	}
	return &fi
}

// The icon has to come first in the markup. By default, the icon will be placed on the left side of the form. To change the alignment, add the flipped option
func (fi *formIcon) IsFlipped() *formIcon {
	fi.isFlipped = true
	return fi
}
func (fi *formIcon) IsClickable() *formIcon {
	fi.isClickable = true
	return fi
}
func (fi *formIcon) Render(ctx context.Context, w io.Writer) error {
	return fi.component().Render(ctx, w)
}

type inputHidden struct {
	component.C
	t     InputType
	id    string
	name  string
	value string
}

func NewInputHidden(name string, attrStyles ...component.AttributesStyles) *inputHidden {
	ih := inputHidden{name: name}
	if len(attrStyles) > 0 {
		ih.SetClasses(attrStyles[0].Styles...)
		if attrStyles[0].Attributes != nil {
			ih.SetAttributes(attrStyles[0].Attributes)
		} else {
			ih.SetAttributes(make(templ.Attributes))
		}
	}
	return &ih
}
func (ih *inputHidden) WithId(id string) *inputHidden {
	ih.id = id
	return ih
}
func (ih *inputHidden) WithValue(val string) *inputHidden {
	ih.value = val
	return ih
}
func (ih *inputHidden) Render(ctx context.Context, w io.Writer) error {
	return ih.component().Render(ctx, w)
}
