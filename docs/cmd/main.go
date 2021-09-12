package main

import (
	"net/url"

	"github.com/hexops/vecty"

	"github.com/mvpmvh/vecty-bootstrap/docs"
)

var (
	AlertPage      docs.Alert
	BadgePage      docs.Badge
	BreadcrumbPage docs.Breadcrumb
	ButtonPage     docs.Button
	router         docs.Router
)

func main() {
	vecty.SetTitle("Vecty Bootstrap Documentation")
	vecty.AddStylesheet(vecty.AddStyleSheetParams{
		URL: "https://cdn.jsdelivr.net/npm/bootstrap-icons@1.5.0/font/bootstrap-icons.css",
	})
	vecty.AddStylesheet(vecty.AddStyleSheetParams{
		URL:         "https://cdn.jsdelivr.net/npm/bootstrap@5.1.0/dist/css/bootstrap.min.css",
		Integrity:   "sha384-KyZXEAg3QhqLMpG8r+8fhAXLRk2vvoC2f3B09zVXn8CA5QIVfZOJ3BCsw2P0p/We",
		CrossOrigin: "anonymous",
	})
	vecty.AddScript(vecty.AddScriptParams{
		URL:         "https://cdn.jsdelivr.net/npm/bootstrap@5.1.0/dist/js/bootstrap.bundle.min.js",
		Integrity:   "sha384-U1DAWAznBHeqEIlVSCgzq+c9gqGAJn5c/t99JyeKa9xxaYpSvHU5awsuZVVFIhvj",
		CrossOrigin: "anonymous",
		AddToHead:   false,
	})

	l := &docs.Layout{Router: &router}

	router.OnRoute = func(r *url.URL) {
		switch r.Fragment {
		case "alerts":
			router.CurrentRoute = &AlertPage
		case "badges":
			router.CurrentRoute = &BadgePage
		case "breadcrumbs":
			router.CurrentRoute = &BreadcrumbPage
		case "buttons":
			router.CurrentRoute = &ButtonPage
		}
		vecty.Rerender(l)
	}

	vecty.RenderBody(l)
}
