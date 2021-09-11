package components

import (
	"fmt"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/mvpmvh/vecty-bootstrap/internal"
)

// SpacingType is either margin or padding
type SpacingType string

const (
	SpacingMargin  SpacingType = "m"
	SpacingPadding SpacingType = "p"
)

// SpacingSide is the location where the margin or padding will be applied
type SpacingSide string

const (
	SideTop       SpacingSide = "t"
	SideBottom    SpacingSide = "b"
	SideStart     SpacingSide = "s"
	SideEnd       SpacingSide = "e"
	SideLeftRight SpacingSide = "x"
	SideTopBottom SpacingSide = "y"
)

// SpacingSize are the sizes margin or padding can be
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

// Spacing provides a container with either margin or padding settings
type Spacing struct {
	vecty.Core
	internal.MarkupData
	SpacingType SpacingType
	SpacingSide SpacingSide
	SpacingSize SpacingSize
	Breakpoint  Breakpoint
	Child       vecty.ComponentOrHTML
}

func (s *Spacing) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			s.MarkupData.Markup(),
			vecty.Class(s.class()),
		),
		s.Child,
	)
}

func (s *Spacing) class() string {
	c := fmt.Sprintf("%s%s", s.SpacingType, s.SpacingSide)
	if s.Breakpoint != "" {
		c = fmt.Sprintf("%s-%s", c, s.Breakpoint)
	}
	return fmt.Sprintf("%s-%s", c, s.SpacingSize)
}
