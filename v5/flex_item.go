package components

import (
	"fmt"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/mvpmvh/vecty-bootstrap/internal"
)

type AlignSelfDirection string

const (
	AlignSelfStart    AlignSelfDirection = "start"
	AlignSelfEnd      AlignSelfDirection = "end"
	AlignSelfCenter   AlignSelfDirection = "center"
	AlignSelfBaseline AlignSelfDirection = "baseline"
	AlignSelfStretch  AlignSelfDirection = "stretch"
)

type FlexItem struct {
	vecty.Core
	internal.MarkupData
	FlexItemSettings
	FlexItemBreakPointSettings []FlexItemBreakpointSettings
	Child                      vecty.ComponentOrHTML
}

type FlexItemSettings struct {
	AlignSelf AlignSelfDirection
	Fill      bool
	Grow      int
	Shrink    int
}

type FlexItemBreakpointSettings struct {
	Breakpoint
	FlexItemSettings
}

func (f *FlexItem) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class(f.classes()...),
			f.MarkupData.Markup(),
		),
		f.Child,
	)
}

func (f *FlexItem) classes() []string {
	var classes []string
	flexItemSettings := append(
		[]FlexItemBreakpointSettings{},
		FlexItemBreakpointSettings{
			FlexItemSettings: f.FlexItemSettings,
		},
	)
	flexItemSettings = append(flexItemSettings, f.FlexItemBreakPointSettings...)
	for _, b := range flexItemSettings {
		if c := f.alignSelfClass(b); c != "" {
			classes = append(classes, c)
		}
		if c := f.fillClass(b); c != "" {
			classes = append(classes, c)
		}
		if c := f.growClass(b); c != "" {
			classes = append(classes, c)
		}
		if c := f.shrinkClass(b); c != "" {
			classes = append(classes, c)
		}
	}

	return classes
}

func (f *FlexItem) alignSelfClass(b FlexItemBreakpointSettings) string {
	if b.AlignSelf == "" {
		return ""
	}

	if b.Breakpoint != "" {
		return fmt.Sprintf("align-self-%s-%s", b.Breakpoint, b.AlignSelf)
	}
	return fmt.Sprintf("align-self-%s", b.AlignSelf)
}

func (f *FlexItem) fillClass(b FlexItemBreakpointSettings) string {
	if !b.Fill {
		return ""
	}

	if b.Breakpoint != "" {
		return fmt.Sprintf("flex-%s-fill", b.Breakpoint)
	}
	return "flex-fill"
}

func (f *FlexItem) growClass(b FlexItemBreakpointSettings) string {
	if b.Grow <= 0 {
		return ""
	}

	if b.Breakpoint != "" {
		return fmt.Sprintf("flex-%s-grow-%d", b.Breakpoint, b.Grow)
	}
	return fmt.Sprintf("flex-grow-%d", b.Grow)
}

func (f *FlexItem) shrinkClass(b FlexItemBreakpointSettings) string {
	if b.Shrink <= 0 {
		return ""
	}

	if b.Breakpoint != "" {
		return fmt.Sprintf("flex-%s-shrink-%d", b.Breakpoint, b.Shrink)
	}
	return fmt.Sprintf("flex-shrink-%d", b.Shrink)
}
