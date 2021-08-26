package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/mvpmvh/vecty-bootstrap/internal"
)

// Row is a row (see:https://getbootstrap.com/docs/5.1/layout/grid/#how-it-works).
type Row struct {
	vecty.Core
	internal.MarkupData
	Child vecty.ComponentOrHTML
}

func (r *Row) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			r.MarkupData.Markup(),
			vecty.Class("row"),
		),
		r.Child,
	)
}
