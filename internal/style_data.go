package internal

import "github.com/hexops/vecty"

// StyleData encapsulates the styles and classes for any component. It should be embedded in every component for the
// component to have a common way to
type StyleData struct {
	Styles  []vecty.Style
	Classes []string
}

// Markup computes the styles and css markup for a component
func (s StyleData) Markup() vecty.MarkupList {
	markup := []vecty.Applyer{
		vecty.Class(s.Classes...),
	}
	for _, style := range s.Styles {
		markup = append(markup, style)
	}

	return vecty.Markup(markup...)
}
