package components

import (
	"fmt"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/mvpmvh/vecty-bootstrap/internal"
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
	internal.StyleData
	Breakpoint ContainerBreakpoint
	Child      vecty.ComponentOrHTML
}

func (c *Container) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			c.StyleData.Markup(),
			vecty.Class(c.class()),
		),
		c.Child,
	)
}

func (c *Container) class() string {
	class := "container"
	if c.Breakpoint != "" {
		class = fmt.Sprintf("container-%s", c.Breakpoint)
	}

	return class
}
