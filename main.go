package main

import (
	"net/http"
	"time"

	"github.com/cspecht/templ_franken-ui/layout/skeleton"
	"github.com/go-chi/chi/v5"
)

func main() {

	r := chi.NewRouter()
	r.Route("/t", func(r chi.Router) {
		r.Get("/", handlePage)

	})
	httpServer := &http.Server{
		// s.cfg.Api.Host +®

		Addr:              "localhost:" + "8080",
		Handler:           r,
		ReadHeaderTimeout: time.Duration(60 * time.Second),
	}
	err := httpServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func handlePage(w http.ResponseWriter, r *http.Request) {
	//page := skeleton.FullSite()

	skeleton.FullSite("name").Render(r.Context(), w)
}
