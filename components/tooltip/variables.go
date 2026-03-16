package tooltip

type Position string
func (p Position) String() string {
	return string(p)
}

const (
	Top Position = "top"
	TopLeft Position = "top-left"
	TopRight Position = "top-right"
	TopCenter Position = "top-center"
	Bottom Position = "bottom"
	BottomLeft Position = "bottom-left"
	BottomRight Position = "bottom-right"
	BottomCenter Position = "bottom-center"
	Left Position = "left"
	Right Position = "right"

)