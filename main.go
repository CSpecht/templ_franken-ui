package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/a-h/templ"
	btn "github.com/cspecht/templ_franken-ui/components/button"
	"github.com/cspecht/templ_franken-ui/components/divider"
	"github.com/cspecht/templ_franken-ui/components/icon"
	"github.com/cspecht/templ_franken-ui/components/label"
	"github.com/cspecht/templ_franken-ui/components/listing"
	"github.com/cspecht/templ_franken-ui/components/navbar"
	"github.com/cspecht/templ_franken-ui/components/placeholder"
	"github.com/cspecht/templ_franken-ui/layout/skeleton"
	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()
	r.Get("/", handlePage)

	fs := http.FileServer(http.Dir("./assets/"))
	r.Handle("/assets/*", http.StripPrefix("/assets", fs))

	httpServer := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadHeaderTimeout: time.Duration(60 * time.Second),
	}
	fmt.Println("Server is running on port 8080")
	err := httpServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	//page := skeleton.FullSite()

	//button
	comp := []templ.Component{}
	comp = append(comp, btn.NewButton("New").Component())
	comp = append(comp, btn.NewButton("Primary").AsPrimary().Component())
	comp = append(comp, btn.NewButton("Icon").SetWidth(btn.W52).SetIcon(icon.NewIcon("copy").SetHeight(20).SetWidth(20).SetStroke(4)).Component())
	comp = append(comp, btn.NewButton("Ghost").AsGhost().Component())
	comp = append(comp, btn.NewButton("Destructive").AsDestructive().Component())
	comp = append(comp, btn.NewButton("Text").AsText().Component())

	//divider
	comp = append(comp, divider.NewDivider().AsIcon().Component())
	comp = append(comp, divider.NewDivider().AsVertical().Component())

	// placeholder
	comp = append(comp, placeholder.NewPlaceholder().Component())

	// label
	l := label.NewLabel("Label")
	l.AsDestructive().AddClasses("test")

	comp = append(comp, l.Component())

	// nav
	nav := navbar.NewNavbar()
	nav.AddItem(nav.AddNavbarEntry("Home").IsActive())
	nav.AddHeader("Header")
	nav.AddDivider()
	nav.AddItem(nav.AddNavbarEntry("Home2").SetLink("#", nil).SetIcon(icon.NewIcon("copy"), nil))
	comp = append(comp, nav.Render())

	//list
	list := listing.NewList().AsStripped()
	list.AddItem(list.AddListEntry("List Entry 1"))
	list.AddItem(list.AddListEntry("List Entry 1"))

	//list.AddItem(list.AddListEntry("List Entry 2"))
	//list.AddItem(list.AddListEntry("List Entry 1"))
	comp = append(comp, list.Render())

	//render fullsite
	skeleton.FullSite("name", comp...).Render(r.Context(), w)
}
