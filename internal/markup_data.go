package internal

import "github.com/hexops/vecty"

// MarkupData encapsulates the styles, classes, and attributes for any component. It should be embedded in every component for the
// component to have a common way to
type MarkupData struct {
	Styles     []vecty.Style
	Classes    []string
	Data       []vecty.Data
	Properties []vecty.Property
}

// Markup computes the styles and css markup for a component
func (s MarkupData) Markup() vecty.MarkupList {
	markup := []vecty.Applyer{
		vecty.Class(s.Classes...),
	}

	for _, style := range s.Styles {
		markup = append(markup, style)
	}

	for _, prop := range s.Properties {
		markup = append(markup, prop)
	}

	for _, data := range s.Data {
		markup = append(markup, data)
	}

	return vecty.Markup(markup...)
}
