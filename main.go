package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/accordion"
	"github.com/cspecht/templ_franken-ui/components/alert"
	"github.com/cspecht/templ_franken-ui/components/breadcrumb"
	"github.com/cspecht/templ_franken-ui/components/button"
	btn "github.com/cspecht/templ_franken-ui/components/button"
	"github.com/cspecht/templ_franken-ui/components/calendar"
	"github.com/cspecht/templ_franken-ui/components/card"
	"github.com/cspecht/templ_franken-ui/components/comment"
	"github.com/cspecht/templ_franken-ui/components/divider"
	"github.com/cspecht/templ_franken-ui/components/dotnav"
	"github.com/cspecht/templ_franken-ui/components/dropdown"
	"github.com/cspecht/templ_franken-ui/components/forms/basics"

	"github.com/cspecht/templ_franken-ui/components/icon"
	"github.com/cspecht/templ_franken-ui/components/image"
	"github.com/cspecht/templ_franken-ui/components/label"
	"github.com/cspecht/templ_franken-ui/components/leader"
	listing "github.com/cspecht/templ_franken-ui/components/list"
	navbar "github.com/cspecht/templ_franken-ui/components/nav"
	"github.com/cspecht/templ_franken-ui/components/pagination"
	"github.com/cspecht/templ_franken-ui/components/placeholder"
	"github.com/cspecht/templ_franken-ui/components/progress"
	"github.com/cspecht/templ_franken-ui/components/spinner"
	"github.com/cspecht/templ_franken-ui/components/table"
	"github.com/cspecht/templ_franken-ui/components/tooltip"
	"github.com/cspecht/templ_franken-ui/components/typography"
	"github.com/cspecht/templ_franken-ui/layout/skeleton"
)

func main() {

	http.HandleFunc("/", handlePage)

	fs := http.FileServer(http.Dir("./assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets", fs))
	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err.Error())
	}

}

func handlePage(w http.ResponseWriter, r *http.Request) {
	//	page := skeleton.FullSite("test", )

	//button
	comp := []templ.Component{}
	comp = append(comp, btn.NewButton("New"))
	comp = append(comp, btn.NewButton("Primary").AsPrimary())
	comp = append(comp, btn.NewButton("Icon").SetWidth40().SetIcon(icon.NewIcon("copy").SetHeight(20).SetWidth(20).SetStroke(4)))
	comp = append(comp, btn.NewButton("Ghost").AsGhost())
	comp = append(comp, btn.NewButton("Destructive").AsDestructive())
	comp = append(comp, btn.NewButton("Text").AsText())

	//divider
	comp = append(comp, divider.NewDivider().WithIcon())

	// placeholder
	comp = append(comp, placeholder.NewPlaceholder())

	// label
	l := label.NewLabel("Label")
	l.AsDestructive().SetClasses("test")

	comp = append(comp, l)

	// nav
	nav := navbar.NewNavigation()
	nav.AsAccordion().AsPrimary()
	nav.AddItem("Home").SetLink("#").SetIcon(icon.NewIcon("copy"))
	nav.AddHeader("Header")
	nav.AddDivider()
	nav.AddItem("Home2").Nav().AddItem("Home3")
	subnav := nav.AddSubNav("More")
	subnav.AddItem("Home5")
	subnav.AddItem("Home6")
	subsubnav := subnav.AddSubNav("SubSubNav")
	subsubnav.AddItem("Home7")
	subsubnav.AddItem("Home8")
	sb, _ := subsubnav.Parent()
	sb.AddItem("Home9")
	sb.Nav().AddItem("Home10")
	comp = append(comp, nav)

	//list
	list := listing.NewList().AsStriped()
	list.AddItem("List Entry 1").List().AddItem("List Entry 2").List().AddItem("List Entry 3")

	//list.AddItem(list.AddListEntry("List Entry 2"))
	//list.AddItem(list.AddListEntry("List Entry 1"))
	comp = append(comp, list.Component())

	comp = append(comp, divider.NewDivider())

	comp = append(comp, alert.NewAlert("Alert").SetTitle("Title").AsCloseable().AsDestructive())
	//render fullsite

	bc := breadcrumb.NewBreadcrumb().AddCrumb("Home", "#").AddCrumb("Home2", "#").AddDisabledCrumb("Home3").AddCurrentCrumb("Home4")

	comp = append(comp, bc)
	cal := calendar.NewCalendar("mycal")
	cal.SetToday()
	comp = append(comp, cal)

	card := card.NewCard()
	card.SetTitle("Card Title").AsDestructive()
	card.SetBody(typography.NewText("This is a card body"))
	card.SetFooter(button.NewButton("Card Footer"))
	comp = append(comp, card)

	// comment
	com := comment.NewComment("Tyler Johnson", "2 days ago", "This is a comment")
	com.SetAvatarImageUrl(templ.SafeURL("https://picsum.photos/200/300"))
	com.SetAvatarAlt("Avatar").AsPrimary()
	comp = append(comp, com)
	com2 := comment.NewComment("Tyler Johnson", "2 days ago", "This is a comment")
	com2.SetAvatarImageUrl(templ.SafeURL("https://picsum.photos/200/300"))
	comlist := comment.NewCommentList()
	comlist.AddComment(com)
	com.AddComment(com2)

	comp = append(comp, comlist)

	// dotnav
	dotnav := dotnav.NewDotnav()
	dotnav.SetVertical()
	dotnav.AddItem("Home").Active()
	dotnav.AddItem("Home2")
	dotnav.AddItem("Home3")
	comp = append(comp, dotnav)

	// accordion
	acc := accordion.NewAccordion()
	acc.AddItem("Item 1", "This is the content of item 1").Accordion().AddItem("Item 2", "This is the content of item 2").IsOpen().Accordion().AddItem("Item 3", "This is the content of item 3")
	comp = append(comp, acc)

	dropDownButton := button.NewButton("Dropdown").AsGhost()
	dropdwn := dropdown.NewDropdown(dropDownButton, nav)
	comp = append(comp, dropdwn)

	// leader
	leader := leader.NewLeader("Chicken Mc Nuggets", "Trölf.Fünfizig")
	leader.SetFillCharacter("#")
	comp = append(comp, leader)
	//scr := scroll.NewScroll("target", Button(Text("Scooll down"))).WithOffset(100)
	//
	img := image.NewImage("https://images.unsplash.com/photo-1495321308589-43affb814eee")
	img.SetClasses("flex", "h-80", "items-center", "justify-center bg-cover")
	img.SetContent(label.NewLabel("some Stuff"))
	comp = append(comp, img)
	progress := progress.NewProgress(10).WithMax(100)
	comp = append(comp, progress)
	spinner := spinner.NewSpinner()
	tltip := tooltip.NewTooltip("Tooltip Title", button.NewButton("tooolitip"))
	comp = append(comp, spinner)
	comp = append(comp, tltip)
	pagi := pagination.NewPagination().WithPreviousLabel("Previous").WithNextLabel("Next")
	pagi.AddPage("1", false, false)
	pagi.AddPage("2", true, false).AddPlacHolder()
	pagi.AddPage("3", false, true)

	comp = append(comp, pagi)
	comp = append(comp, list)
	tabelle := table.NewTable()
	tabelle.WithCaption("this is a wonderful table").IsStriped().IsResponsiveStack()
	tabelle.AddHeader(tabelle.NewCellElementText("ID"), tabelle.NewCellElementText("Username"), tabelle.NewCellElementText("Email"))
	tabelle.NewRowElement().AddCells(tabelle.NewCellElementText("123"), tabelle.NewCellElementText("peter"), tabelle.NewCellElementText("test@gmail.com"))
	tabelle.NewRowElement().AddCells(tabelle.NewCellElementText("123"), tabelle.NewCellElementText("peter"), tabelle.NewCellElementText("test@gmail.com"))
	comp = append(comp, tabelle)
	f := basics.ExampleForm()
	comp = append(comp, f)
	f2 := basics.AnotherExample()
	comp = append(comp, f2)
	form := basics.NewForm()
	fs := basics.NewFieldSet().WithLegend("My User Inputs").AddComponents(basics.NewLabel("Best","best"),basics.NewFormControls(basics.NewInputText("name","","").DisableSpellcheck().DisableAutocomplete()),basics.NewErrorText("super error", "bla"))
	form.AddComponents(fs)
	form.AddActions(button.NewButton("Submit"))
	comp = append(comp, form)
	templ.Handler(skeleton.FullSite("name", comp...)).ServeHTTP(w, r)
	//scr.Render(w)

}
