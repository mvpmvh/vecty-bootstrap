package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

// Row is a row (see:https://getbootstrap.com/docs/5.1/layout/grid/#how-it-works).
type Row struct {
	vecty.Core
	Styles  []vecty.Style
	Classes []string
	Child   vecty.ComponentOrHTML
}

func (r *Row) Render() vecty.ComponentOrHTML {
	return elem.Div(
		r.markup(),
		r.Child,
	)
}

func (r *Row) markup() vecty.MarkupList {
	markup := []vecty.Applyer{
		vecty.Class(vecty.AppendClasses([]string{"row"}, r.Classes)...),
	}
	for _, style := range r.Styles {
		markup = append(markup, style)
	}

	return vecty.Markup(markup...)
}
