package components

import (
	"fmt"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/mvpmvh/vecty-bootstrap/internal"
)

// ColumnBreakpoint is a responsive breakpoint for a Column
type ColumnBreakpoint string

// The different breakpoints for a row
const (
	ColumnSmall           ColumnBreakpoint = "sm"
	ColumnMedium          ColumnBreakpoint = "md"
	ColumnLarge           ColumnBreakpoint = "lg"
	ColumnExtraLarge      ColumnBreakpoint = "xl"
	ColumnExtraExtraLarge ColumnBreakpoint = "xxl"
	ColumnAuto            ColumnBreakpoint = "auto"
)

// ColumnBreakpointSize configures a Column's responsive breakpoint and size settings
type ColumnBreakpointSize struct {
	Breakpoint
	SpacingSize
}

// Column is a Bootstrap Column (see: https://getbootstrap.com/docs/5.1/layout/grid/#how-it-works)
type Column struct {
	vecty.Core
	internal.MarkupData
	Breakpoints []ColumnBreakpointSize
	Child       vecty.ComponentOrHTML
}

func (c *Column) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class(c.classes()...),
			c.MarkupData.Markup(),
		),
		c.Child,
	)
}

func (c *Column) classes() []string {
	var classes []string
	for _, b := range c.Breakpoints {
		class := "col"
		if b.Breakpoint != "" {
			class = fmt.Sprintf("col-%s", b.Breakpoint)
		}
		if b.SpacingSize != "" {
			class = fmt.Sprintf("%s-%s", class, b.SpacingSize)
		}
		classes = append(classes, class)
	}

	if classes == nil {
		classes = []string{"col"}
	}

	return classes
}
