package button

type buttonVariant string

const (
	Default     buttonVariant = "uk-btn-default"
	Ghost       buttonVariant = "uk-btn-ghost"
	Primary     buttonVariant = "uk-btn-primary"
	Secondary   buttonVariant = "uk-btn-secondary"
	Destructive buttonVariant = "uk-btn-destructive"
	Text        buttonVariant = "uk-btn-text"
	Link        buttonVariant = "uk-btn-link"
)

type buttonSize string

const (
	Xs buttonSize = "uk-btn-xs"
	Sm buttonSize = "uk-btn-sm"
	Md buttonSize = "uk-btn-md"
	Lg buttonSize = "uk-btn-lg"
)

type withModifiers string

const (
	W40   withModifiers = "w-40"
	W44   withModifiers = "w-44"
	W48   withModifiers = "w-48"
	W52   withModifiers = "w-52"
	WFull withModifiers = "w-full"
)
