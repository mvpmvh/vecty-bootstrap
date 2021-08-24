package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

// AnchorButton is an <a> tag that appears as as <button>. It has an href set to `#` and role="button" (see: https://getbootstrap.com/docs/5.1/components/buttons/#button-tags)
type AnchorButton struct {
	Button
}

func (a *AnchorButton) Render() vecty.ComponentOrHTML {
	markup := []vecty.Applyer{prop.Href("#")}
	return elem.Anchor(
		vecty.MergeMarkupList(a.markup(), vecty.Markup(markup...)),
		a.Child,
	)
}
