package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/mvpmvh/vecty-bootstrap/internal"
)

// Collapse provides collapsible content (see: https://getbootstrap.com/docs/5.1/components/collapse/)
type Collapse struct {
	vecty.Core
	internal.MarkupData
	IsShowing    bool
	IsHorizontal bool
	Child        vecty.ComponentOrHTML
}

func (c *Collapse) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("collapse"),
			c.MarkupData.Markup(),
			vecty.ClassMap{
				"show":                c.IsShowing,
				"collapse-horizontal": c.IsHorizontal,
			},
		),
		c.Child,
	)
}
