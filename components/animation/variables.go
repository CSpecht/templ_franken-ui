package animation

import "github.com/a-h/templ"

// These are the animation types and their corresponding CSS classes.
const(
	FadeIn string = "uk-anmt-fade"
	ScaleUp string = "uk-anmt-scale-up"
	ScaleDown string = "uk-anmt-scale-down"
	Shake string = "uk-anmt-shake"
	SlideLeft string = "uk-anmt-slide-left"
	SlideTop string = "uk-anmt-slide-top"
	SlideBottom string = "uk-anmt-slide-bottom"
	SlideRight string = "uk-anmt-slide-right"
	SlideLeftSmall string = "uk-anmt-slide-left-sm"
	SlideTopSmall string = "uk-anmt-slide-top-sm"
	SlideBottomSmall string = "uk-anmt-slide-bottom-sm"
	SlideRightSmall string = "uk-anmt-slide-right-sm"
	SlideLeftMedium string = "uk-anmt-slide-left-md"
	SlideTopMedium string = "uk-anmt-slide-top-md"
	SlideBottomMedium string = "uk-anmt-slide-bottom-md"
	SlideRightMedium string = "uk-anmt-slide-right-md"
	ReverseModifier string = "uk-anmt-reverse" // By default, all animations are incoming. To reverse any animation, add this class.
	FastModifier string = "uk-anmt-fast"
	KenBursn string = "uk-anmt-kenburns"
	SVGStroke string = "uk-anmt-svg-stroke" // This animation is only available for SVG elements. It animates the stroke of an SVG path.
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