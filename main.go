package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/cspecht/templ_franken-ui/components/alert"
	"github.com/cspecht/templ_franken-ui/components/breadcrumb"
	btn "github.com/cspecht/templ_franken-ui/components/button"
	"github.com/cspecht/templ_franken-ui/components/divider"
	"github.com/cspecht/templ_franken-ui/components/icon"
	"github.com/cspecht/templ_franken-ui/components/label"
	listing "github.com/cspecht/templ_franken-ui/components/list"
	navbar "github.com/cspecht/templ_franken-ui/components/nav"
	"github.com/cspecht/templ_franken-ui/components/placeholder"
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
	sb,_ := subsubnav.Parent()
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

	templ.Handler(skeleton.FullSite("name", comp...)).ServeHTTP(w, r)

}
