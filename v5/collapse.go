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
	ID           string
	IsShowing    bool
	IsHorizontal bool
	Child        vecty.ComponentOrHTML
}

func (c *Collapse) Render() vecty.ComponentOrHTML {
	if c.ID == "" {
		panic("no id set.")
	}

	return elem.Div(
		vecty.Markup(
			c.MarkupData.Markup(),
			vecty.ID(c.ID),
			vecty.Class("collapse"),
			vecty.ClassMap{
				"show":                c.IsShowing,
				"collapse-horizontal": c.IsHorizontal,
			},
		),
		c.Child,
	)
}
