package basics

type BasicTypes string

func (b BasicTypes) String() string { return string(b) }

const (
	Input        BasicTypes = "uk-input"
	Select       BasicTypes = "uk-select"
	TextArea     BasicTypes = "uk-textarea"
	Radio        BasicTypes = "uk-radio"
	Checkbox     BasicTypes = "uk-checkbox"
	Range        BasicTypes = "uk-range"
	ToggleSwitch BasicTypes = "uk-toggle-switch"
)

type InputType string

func (i InputType) String() string { return string(i) }

const (
	Text   InputType = "text"
	Hidden InputType = "hidden"
	Number InputType = "number"
)

type SizeModifier string

func (sm SizeModifier) String() string { return string(sm) }

const (
	XsSize SizeModifier = "uk-form-xs"
	SmSize SizeModifier = "uk-form-sm"
	MdSize SizeModifier = "uk-form-md"
	LgSize SizeModifier = "uk-form-lg"
)
type FormLayout string

func (fl FormLayout) String() string { return string(fl) }

const (
	Stacked FormLayout = "uk-form-stacked"
	Horizontal FormLayout = "uk-form-horizontal"
)