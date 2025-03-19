package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/cspecht/templ_franken-ui/layout/skeleton"
	btn "github.com/cspecht/templ_franken-ui/components/button"
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

	skeleton.FullSite("name", btn.NewDefaultButton()).Render(r.Context(), w)
}
