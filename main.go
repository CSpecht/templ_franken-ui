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
	"github.com/cspecht/templ_franken-ui/components/icon"
	"github.com/cspecht/templ_franken-ui/components/label"
	"github.com/cspecht/templ_franken-ui/components/leader"
	listing "github.com/cspecht/templ_franken-ui/components/list"
	navbar "github.com/cspecht/templ_franken-ui/components/nav"
	"github.com/cspecht/templ_franken-ui/components/placeholder"
	"github.com/cspecht/templ_franken-ui/components/scroll"
	"github.com/cspecht/templ_franken-ui/components/typography"

	. "maragu.dev/gomponents"

	. "maragu.dev/gomponents/html"
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
	//page := skeleton.FullSite()

	//button
	comp := []templ.Component{}
	comp = append(comp, btn.NewButton("New").Component())
	comp = append(comp, btn.NewButton("Primary").AsPrimary().Component())
	comp = append(comp, btn.NewButton("Icon").SetWidth40().SetIcon(icon.NewIcon("copy").SetHeight(20).SetWidth(20).SetStroke(4)).Component())
	comp = append(comp, btn.NewButton("Ghost").AsGhost().Component())
	comp = append(comp, btn.NewButton("Destructive").AsDestructive().Component())
	comp = append(comp, btn.NewButton("Text").AsText().Component())

	//divider
	comp = append(comp, divider.NewDivider().WithIcon().Component())

	// placeholder
	comp = append(comp, placeholder.NewPlaceholder().Component())

	// label
	l := label.NewLabel("Label")
	l.AsDestructive().AddClasses("test")

	comp = append(comp, l.Component())

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
	comp = append(comp, nav.Component())

	//list
	list := listing.NewList().AsStriped()
	list.AddItem("List Entry 1").List().AddItem("List Entry 2").List().AddItem("List Entry 3")

	//list.AddItem(list.AddListEntry("List Entry 2"))
	//list.AddItem(list.AddListEntry("List Entry 1"))
	comp = append(comp, list.Component())

	comp = append(comp, divider.NewDivider().Component())

	comp = append(comp, alert.NewAlert("Alert").SetTitle("Title").AsCloseable().AsDestructive().Component())
	//render fullsite

	bc := breadcrumb.NewBreadcrumb().AddCrumb("Home", "#").AddCrumb("Home2", "#").AddDisabledCrumb("Home3").AddCurrentCrumb("Home4")

	comp = append(comp, bc.Component())
	cal := calendar.NewCalendar("mycal")
	cal.SetToday()
	comp = append(comp, cal.Component())

	card := card.NewCard()
	card.SetTitle("Card Title").AsDestructive()
	card.SetBody(typography.NewText("This is a card body").Component())
	card.SetFooter(button.NewButton("Card Footer").Component())
	comp = append(comp, card.Component())

	// comment
	com := comment.NewComment("Tyler Johnson", "2 days ago", "This is a comment")
	com.SetAvatarImageUrl(templ.SafeURL("https://picsum.photos/200/300"))
	com.SetAvatarAlt("Avatar").AsPrimary()
	comp = append(comp, com.Component())
	com2 := comment.NewComment("Tyler Johnson", "2 days ago", "This is a comment")
	com2.SetAvatarImageUrl(templ.SafeURL("https://picsum.photos/200/300"))
	comlist := comment.NewCommentList()
	comlist.AddComment(com)
	com.AddComment(com2)

	comp = append(comp, comlist.Component())

	// dotnav
	dotnav := dotnav.NewDotnav()
	dotnav.SetVertical()
	dotnav.AddItem("Home").Active()
	dotnav.AddItem("Home2")
	dotnav.AddItem("Home3")
	comp = append(comp, dotnav.Component())

	// accordion
	acc := accordion.NewAccordion()
	acc.AddItem("Item 1", "This is the content of item 1").Accordion().AddItem("Item 2", "This is the content of item 2").IsOpen().Accordion().AddItem("Item 3", "This is the content of item 3")
	comp = append(comp, acc.Component())

	dropDownButton := button.NewButton("Dropdown").AsPrimary()
	dropdwn := dropdown.NewDropdown(dropDownButton.Component(), nav)
	comp = append(comp, dropdwn.Component())

	// leader
	leader := leader.NewLeader("Chicken Mc Nuggets", "Trölf.Fünfizig")
	leader.SetFillCharacter("#")
	comp = append(comp, leader.Component())
	scr := scroll.NewScroll("target", Button(Text("Scooll down"))).WithOffset(100)

	//templ.Handler(skeleton.FullSite("name", comp...)).ServeHTTP(w, r)
	scr.Render(w)

}
