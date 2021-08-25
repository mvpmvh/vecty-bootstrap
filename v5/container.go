package components

import (
	"fmt"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

// ContainerBreakpoint is a breakpoint for responsive container sizes
type ContainerBreakpoint string

const (
	ContainerSmall           ContainerBreakpoint = "sm"
	ContainerMedium          ContainerBreakpoint = "md"
	ContainerLarge           ContainerBreakpoint = "lg"
	ContainerExtraLarge      ContainerBreakpoint = "xl"
	ContainerExtraExtraLarge ContainerBreakpoint = "xxl"
	ContainerFluid           ContainerBreakpoint = "fluid"
)

// Container is a Bootstrap container (see: https://getbootstrap.com/docs/5.1/layout/containers/).
type Container struct {
	vecty.Core
	Styles     []vecty.Style
	Classes    []string
	Breakpoint ContainerBreakpoint
	Child      vecty.ComponentOrHTML
}

func (c *Container) Render() vecty.ComponentOrHTML {
	return elem.Div(
		c.markup(),
		vecty.If(c.Child != nil, c.Child),
	)
}

func (c *Container) markup() vecty.MarkupList {
	class := "container"
	if c.Breakpoint != "" {
		class = fmt.Sprintf("container-%s", c.Breakpoint)
	}

	classes := vecty.AppendClasses([]string{class}, c.Classes)
	markup := []vecty.Applyer{
		vecty.Class(classes...),
	}

	for _, style := range c.Styles {
		markup = append(markup, style)
	}

	return vecty.Markup(markup...)
}
