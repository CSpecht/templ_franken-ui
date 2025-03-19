package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/a-h/templ"
	btn "github.com/cspecht/templ_franken-ui/components/button"
	"github.com/cspecht/templ_franken-ui/components/icon"
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
	buttons := []templ.Component{}
	buttons = append(buttons, btn.NewButton("New").Component())
	buttons = append(buttons, btn.NewButton("Primary").AsPrimary().Component())
	buttons = append(buttons, btn.NewButton("Icon").SetWidth(btn.W52).SetIcon(icon.NewIcon("copy").SetHeight(20).SetWidth(20).SetStroke(4)).Component())
	buttons = append(buttons, btn.NewButton("Ghost").AsGhost().Component())
	buttons = append(buttons, btn.NewButton("Destructive").AsDestructive().Component())
	buttons = append(buttons, btn.NewButton("Text").AsText().Component())
	//render fullsite
	skeleton.FullSite("name", buttons...).Render(r.Context(), w)

}
