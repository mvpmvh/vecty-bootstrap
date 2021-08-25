package components

import (
	"fmt"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type SpacingType string

const (
	SpacingMargin  SpacingType = "m"
	SpacingPadding SpacingType = "p"
)

type SpacingSide string

const (
	SideTop       SpacingSide = "t"
	SideBottom    SpacingSide = "b"
	SideStart     SpacingSide = "s"
	SideEnd       SpacingSide = "e"
	SideLeftRight SpacingSide = "x"
	SideTopBottom SpacingSide = "y"
)

type SpacingSize string

const (
	SizeZero  SpacingSize = "0"
	SizeOne   SpacingSize = "1"
	SizeTwo   SpacingSize = "2"
	SizeThree SpacingSize = "3"
	SizeFour  SpacingSize = "4"
	SizeFive  SpacingSize = "5"
	SizeAuto  SpacingSize = "auto"
)

type Spacing struct {
	vecty.Core

	SpacingType SpacingType
	SpacingSide SpacingSide
	SpacingSize SpacingSize
	Breakpoint  Breakpoint
	Children    vecty.List
}

func (s *Spacing) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class(s.class()),
		),
		s.Children,
	)
}

func (s *Spacing) class() string {
	c := fmt.Sprintf("%s%s", s.SpacingType, s.SpacingSide)
	if s.Breakpoint != "" {
		c = fmt.Sprintf("%s-%s", c, s.Breakpoint)
	}
	return fmt.Sprintf("%s-%s", c, s.SpacingSize)
}
