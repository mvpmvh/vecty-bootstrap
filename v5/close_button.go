package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/mvpmvh/vecty-bootstrap/internal"
)

// CloseButton is a generic close button for dismissing content like modals and alerts (see: https://getbootstrap.com/docs/5.1/components/close-button/)
type CloseButton struct {
	vecty.Core
	internal.MarkupData
	IsDisabled bool
	IsWhite    bool
}

func (c *CloseButton) Render() vecty.ComponentOrHTML {
	return elem.Button(
		vecty.Markup(
			vecty.Class("btn-close"),
			vecty.ClassMap{"btn-close-white": c.IsWhite},
			vecty.Type(vecty.TypeButton),
			vecty.Property{"aria-label": "Close"},
			vecty.MarkupIf(c.IsDisabled, vecty.Disabled(true)),
		),
	)
}
