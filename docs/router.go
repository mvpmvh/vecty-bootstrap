package docs

import (
	"fmt"
	"net/url"

	"github.com/hexops/vecty"
)

// Router notifies a listen of route changes
type Router struct {
	OnRoute      func(*url.URL)
	CurrentRoute vecty.ComponentOrHTML
}

// To updates the browser history with the newest URL to visit. It notifies its OnRoute listener.
func (r *Router) To(s string) {
	u, err := url.Parse(s)
	if err != nil {
		fmt.Println("error parsing url: " + err.Error())
		return
	}

	// TODO: push path to window.history
	if r.OnRoute != nil {
		r.OnRoute(u)
	}
}
