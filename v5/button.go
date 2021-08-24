package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/prop"
)

// Button creates a button (see: https://getbootstrap.com/docs/5.1/components/buttons/)
type Button struct {
	vecty.Core
	Styles     []vecty.Style
	Classes    []string
	Type       prop.InputType
	Color      Color
	Size       ButtonSize
	IsDisabled bool
	IsOutline  bool
	IsLink     bool
	Child      vecty.ComponentOrHTML
}

// ButtonSize controls the size of the button (see: https://getbootstrap.com/docs/5.1/components/buttons/#sizes)
type ButtonSize int

const (
	UnknownSize ButtonSize = iota
	SmallSize
	LargeSize
)

func (b *Button) Render() vecty.ComponentOrHTML {
	return elem.Button(
		b.markup(),
		b.Child,
	)
}

func (b *Button) markup() vecty.MarkupList {
	classes := []string{"btn"}
	classes = append(classes, b.Classes...)
	markup := []vecty.Applyer{
		vecty.Class(classes...),
		b.classMap(),
		prop.Type(b.Type),
		prop.Disabled(b.IsDisabled),
	}

	for _, style := range b.Styles {
		markup = append(markup, style)
	}

	return vecty.Markup(markup...)
}

func (b *Button) classMap() vecty.ClassMap {
	return vecty.ClassMap{
		"btn-lg":                b.Size == LargeSize,
		"btn-sm":                b.Size == SmallSize,
		"btn-outline-primary":   b.IsOutline && b.Color == PrimaryColor,
		"btn-outline-secondary": b.IsOutline && b.Color == SecondaryColor,
		"btn-outline-success":   b.IsOutline && b.Color == SuccessColor,
		"btn-outline-danger":    b.IsOutline && b.Color == DangerColor,
		"btn-outline-warning":   b.IsOutline && b.Color == WarningColor,
		"btn-outline-info":      b.IsOutline && b.Color == InfoColor,
		"btn-outline-light":     b.IsOutline && b.Color == LightColor,
		"btn-outline-dark":      b.IsOutline && b.Color == DarkColor,
		"btn-primary":           !b.IsOutline && b.Color == PrimaryColor,
		"btn-secondary":         !b.IsOutline && b.Color == SecondaryColor,
		"btn-success":           !b.IsOutline && b.Color == SuccessColor,
		"btn-danger":            !b.IsOutline && b.Color == DangerColor,
		"btn-warning":           !b.IsOutline && b.Color == WarningColor,
		"btn-info":              !b.IsOutline && b.Color == InfoColor,
		"btn-light":             !b.IsOutline && b.Color == LightColor,
		"btn-dark":              !b.IsOutline && b.Color == DarkColor,
		"btn-link":              b.IsLink,
	}
}
