package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/mvpmvh/vecty-bootstrap/internal"
)

// Badge creates an alert (see: https://getbootstrap.com/docs/5.1/components/badge/)
type Badge struct {
	vecty.Core
	internal.MarkupData
	Child vecty.ComponentOrHTML
}

func (b *Badge) Render() vecty.ComponentOrHTML {
	return elem.Span(
		vecty.Markup(
			b.MarkupData.Markup(),
			vecty.Class("badge"),
		),
		b.Child,
	)
}
