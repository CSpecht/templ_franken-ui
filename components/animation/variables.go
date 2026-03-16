package animation

import "github.com/a-h/templ"

type Animation string
func (a Animation) String() string {
	return string(a)
}
// These are the animation types and their corresponding CSS classes.
const(
	FadeIn Animation = "uk-anmt-fade"
	ScaleUp Animation = "uk-anmt-scale-up"
	ScaleDown Animation = "uk-anmt-scale-down"
	Shake Animation = "uk-anmt-shake"
	SlideLeft Animation = "uk-anmt-slide-left"
	SlideTop Animation = "uk-anmt-slide-top"
	SlideBottom Animation = "uk-anmt-slide-bottom"
	SlideRight Animation = "uk-anmt-slide-right"
	SlideLeftSmall Animation = "uk-anmt-slide-left-sm"
	SlideTopSmall Animation = "uk-anmt-slide-top-sm"
	SlideBottomSmall Animation = "uk-anmt-slide-bottom-sm"
	SlideRightSmall Animation = "uk-anmt-slide-right-sm"
	SlideLeftMedium Animation = "uk-anmt-slide-left-md"
	SlideTopMedium Animation = "uk-anmt-slide-top-md"
	SlideBottomMedium Animation = "uk-anmt-slide-bottom-md"
	SlideRightMedium Animation = "uk-anmt-slide-right-md"
	ReverseModifier Animation = "uk-anmt-reverse" // By default, all animations are incoming. To reverse any animation, add this class.
	FastModifier Animation = "uk-anmt-fast"
	KenBursn Animation = "uk-anmt-kenburns"
	SVGStroke Animation = "uk-anmt-svg-stroke" // This animation is only available for SVG elements. It animates the stroke of an SVG path.
)
// By default, scaling animations originate from the center. To modify this behavior add this class with the desired origin.
// The origin can be any of the following: top-left, top-right, etc .
func OriginModifier(first, second string) string {
	return "uk-anmt-origin-" + first + "-" + second
}

// The Animation component can be used to animate SVG strokes. The effect looks like the SVG strokes are drawn before your eyes. 
// This needs to be added to the img tag of the SVG element.
func SVGStrokes() templ.Attributes{
	attr := make(templ.Attributes)
	attr["data-uk-svg"] = "stroke-animation: true"
	return attr
}