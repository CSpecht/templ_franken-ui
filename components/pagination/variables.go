package pagination

type PaginationState string

func (s PaginationState) String() string {
	return string(s)
}

const (
	Active   PaginationState = "uk-active"
	Disabled PaginationState = "uk-disabled"
)

type PaginationStyle string

func (s PaginationStyle) String() string {
	return string(s)
}

const (
	DefaultStyle PaginationStyle = "uk-pgn-default"
	Primary      PaginationStyle = "uk-pgn-primary"
	Secondary    PaginationStyle = "uk-pgn-secondary"
	Destructive  PaginationStyle = "uk-pgn-destructive"
	Ghost        PaginationStyle = "uk-pgn-ghost"
)

type PaginationSize string

func (s PaginationSize) String() string {
	return string(s)
}

const (
	Xs PaginationSize = "uk-pgn-xs"
	Sm PaginationSize = "uk-pgn-sm"
	Md PaginationSize = "uk-pgn-md"
	Lg PaginationSize = "uk-pgn-lg"
)

// The Pagination component utilizes flexbox, so navigations can easily be aligned with Flex utility classes from Tailwind CSS
type PaginationAlignment string

func (a PaginationAlignment) String() string {
	return string(a)
}

const (
	DefaultAlignment PaginationAlignment = ""
	Center           PaginationAlignment = "justify-center"
	Left             PaginationAlignment = "justify-start"
	Right            PaginationAlignment = "justify-end"
	Between          PaginationAlignment = "justify-between"
)
