package components

import (
	"fmt"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/mvpmvh/vecty-bootstrap/internal"
)

type FlexOrder string

const (
	Order0     FlexOrder = "0"
	Order1     FlexOrder = "1"
	Order2     FlexOrder = "2"
	Order3     FlexOrder = "3"
	Order4     FlexOrder = "4"
	Order5     FlexOrder = "5"
	OrderFirst FlexOrder = "first"
	OrderLast  FlexOrder = "last"
)

type FlexWrap string

const (
	WrapNone     FlexWrap = "nowrap"
	WrapForward  FlexWrap = "wrap"
	WrapReversed FlexWrap = "wrap-reverse"
)

type AlignContentDirection string

const (
	AlignContentStart   AlignContentDirection = "start"
	AlignContentEnd     AlignContentDirection = "end"
	AlignContentCenter  AlignContentDirection = "center"
	AlignContentBetween AlignContentDirection = "between"
	AlignContentAround  AlignContentDirection = "around"
	AlignContentStretch AlignContentDirection = "stretch"
)

type AlignItemsDirection string

const (
	AlignItemsStart    AlignItemsDirection = "start"
	AlignItemsEnd      AlignItemsDirection = "end"
	AlignItemsCenter   AlignItemsDirection = "center"
	AlignItemsBaseline AlignItemsDirection = "baseline"
	AlignItemsStretch  AlignItemsDirection = "stretch"
)

type FlexJustifyDirection string

const (
	JustifyContentStart   FlexJustifyDirection = "start"
	JustifyContentEnd     FlexJustifyDirection = "end"
	JustifyContentCenter  FlexJustifyDirection = "center"
	JustifyContentBetween FlexJustifyDirection = "between"
	JustifyContentAround  FlexJustifyDirection = "between"
)

type FlexBreakpointSettings struct {
	Breakpoint
	FlexSettings
}

type FlexSettings struct {
	IsInline       bool
	IsVertical     bool
	IsReversed     bool
	Wrap           FlexWrap
	JustifyContent FlexJustifyDirection
	AlignItems     AlignItemsDirection
	AlignContent   AlignContentDirection
	Order          FlexOrder
}

type Flex struct {
	vecty.Core
	internal.MarkupData
	FlexSettings
	FlexBreakpointSettings []FlexBreakpointSettings
	Children               vecty.List
}

func (f *Flex) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			f.MarkupData.Markup(),
			vecty.Class(f.classes()...),
		),
		f.Children,
	)
}

func (f *Flex) classes() []string {
	classes := append([]string{}, f.Classes...)
	flexBreakpointSettings := append(
		[]FlexBreakpointSettings{},
		FlexBreakpointSettings{
			FlexSettings: f.FlexSettings,
		},
	)
	flexBreakpointSettings = append(flexBreakpointSettings, f.FlexBreakpointSettings...)
	for _, b := range flexBreakpointSettings {
		classes = append(classes, f.displayClass(b), f.directionClass(b))
		if c := f.justifyContentClass(b); c != "" {
			classes = append(classes, c)
		}
		if c := f.alignItemsClass(b); c != "" {
			classes = append(classes, c)
		}
		if c := f.alignContentClass(b); c != "" {
			classes = append(classes, c)
		}
		if c := f.wrapClass(b); c != "" {
			classes = append(classes, c)
		}
		if c := f.orderClass(b); c != "" {
			classes = append(classes, c)
		}
	}

	return classes
}

func (f *Flex) displayClass(settings FlexBreakpointSettings) string {
	if settings.IsInline {
		if settings.Breakpoint != "" {
			return fmt.Sprintf("d-%s-inline-flex", settings.Breakpoint)
		}
		return "d-inline-flex"
	}

	if settings.Breakpoint != "" {
		return fmt.Sprintf("d-%s-flex", settings.Breakpoint)
	}

	return "d-flex"
}

func (f *Flex) directionClass(settings FlexBreakpointSettings) string {
	c := "flex"
	dir := "row"
	if settings.IsVertical {
		dir = "col"
	}

	c = fmt.Sprintf("%s-%s", c, dir)
	if settings.Breakpoint != "" {
		c = fmt.Sprintf("%s-%s", c, settings.Breakpoint)
	}

	if settings.IsReversed {
		c = fmt.Sprintf("%s-reverse", c)
	}

	return c
}

func (f *Flex) justifyContentClass(settings FlexBreakpointSettings) string {
	if settings.JustifyContent == "" {
		return ""
	}

	if settings.Breakpoint != "" {
		return fmt.Sprintf("justify-content-%s-%s", settings.Breakpoint, settings.JustifyContent)
	}

	return fmt.Sprintf("justify-content-%s", settings.JustifyContent)
}

func (f *Flex) alignItemsClass(settings FlexBreakpointSettings) string {
	if settings.AlignItems == "" {
		return ""
	}

	if settings.Breakpoint != "" {
		return fmt.Sprintf("align-items-%s-%s", settings.Breakpoint, settings.AlignItems)
	}
	return fmt.Sprintf("align-items-%s", settings.AlignItems)
}

func (f *Flex) alignContentClass(settings FlexBreakpointSettings) string {
	if settings.AlignContent == "" {
		return ""
	}

	if settings.Breakpoint != "" {
		return fmt.Sprintf("align-content-%s-%s", settings.Breakpoint, settings.AlignContent)
	}
	return fmt.Sprintf("align-content-%s", settings.AlignContent)
}

func (f *Flex) wrapClass(settings FlexBreakpointSettings) string {
	if settings.Wrap == "" {
		return ""
	}

	if settings.Breakpoint != "" {
		return fmt.Sprintf("flex-wrap-%s-%s", settings.Breakpoint, settings.Wrap)
	}
	return fmt.Sprintf("flex-wrap-%s", settings.Wrap)
}

func (f *Flex) orderClass(settings FlexBreakpointSettings) string {
	if settings.Order == "" {
		return ""
	}

	if settings.Breakpoint != "" {
		return fmt.Sprintf("order-%s-%s", settings.Breakpoint, settings.Order)
	}
	return fmt.Sprintf("flex-wrap-%s", settings.Order)
}
