package button

type variant string

const (
	Default     variant = "uk-btn-default"
	Ghost       variant = "uk-btn-ghost"
	Primary     variant = "uk-btn-primary"
	Secondary   variant = "uk-btn-secondary"
	Destructive variant = "uk-btn-destructive"
	Text        variant = "uk-btn-text"
	Link        variant = "uk-btn-link"
)

type size string

const (
	Xs size = "uk-btn-xs"
	Sm size = "uk-btn-sm"
	Md size = "uk-btn-md"
	Lg size = "uk-btn-lg"
)

type width string

const (
	W40   width = "w-40"
	W44   width = "w-44"
	W48   width = "w-48"
	W52   width = "w-52"
	WFull width = "w-full"
)
