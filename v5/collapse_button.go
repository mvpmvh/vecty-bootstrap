package components

import (
	"github.com/hexops/vecty"

	"github.com/mvpmvh/vecty-bootstrap/internal"
)

// CollapseButton is a button which toggles collapsible content (see: https://getbootstrap.com/docs/5.1/components/collapse/)
type CollapseButton struct {
	vecty.Core
	Button
	// Target is the css selector to target as being collapsible
	Target string
}

func (c *CollapseButton) Render() vecty.ComponentOrHTML {
	data := append(
		[]vecty.Data{
			{
				"bsToggle": "collapse",
				"bsTarget": c.Target,
			},
		},
		c.Button.MarkupData.Data...,
	)

	b := Button{
		MarkupData: internal.MarkupData{
			Styles:     c.Button.MarkupData.Styles,
			Classes:    c.Button.MarkupData.Classes,
			Properties: c.Button.MarkupData.Properties,
			Data:       data,
		},
		Type:       vecty.TypeButton,
		Color:      c.Button.Color,
		Size:       c.Button.Size,
		IsDisabled: c.Button.IsDisabled,
		IsOutline:  c.Button.IsOutline,
		IsLink:     c.Button.IsLink,
		Child:      c.Button.Child,
	}

	return b.Render()
}
