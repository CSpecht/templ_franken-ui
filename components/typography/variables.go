package typography

type HeadingSize string

const (
	H1 HeadingSize = "uk-h1"
	H2 HeadingSize = "uk-h2"
	H3 HeadingSize = "uk-h3"
	H4 HeadingSize = "uk-h4"
)

type HeadingStyle string

const (
	Divider HeadingStyle = "uk-heading-divider"
	Line    HeadingStyle = "uk-heading-line"
	Bullet  HeadingStyle = "uk-heading-bullet"
)

type HeroSize string

const (
	SM      HeroSize = "uk-hero-sm"
	MD      HeroSize = "uk-hero-md"
	LG      HeroSize = "uk-hero-lg"
	XL      HeroSize = "uk-hero-xl"
	TwoXL   HeroSize = "uk-hero-2xl"
	ThreeXL HeroSize = "uk-hero-3xl"
)

type Alignment string

const (
	Left   Alignment = "uk-text-left"
	Center Alignment = "text-center"
	Right  Alignment = "uk-text-right"
)

type TextSize string
const (
	Small TextSize = "uk-text-sm"
	Medium TextSize = "uk-text-md"
	Large TextSize = "uk-text-lg"
	Default TextSize = "uk-text-base"
)

type TextStyle string
const (
	Lead TextStyle = "uk-text-lead"
	Meta TextStyle = "uk-text-meta"
	Truncate TextStyle = "uk-text-truncate"
	Break TextStyle = "uk-text-break"
)