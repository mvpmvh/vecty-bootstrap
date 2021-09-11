package components

import (
	"fmt"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/mvpmvh/vecty-bootstrap/internal"
)

// Carousel creates a carousel (see: https://getbootstrap.com/docs/5.1/components/carousel/)
type Carousel struct {
	vecty.Core
	internal.MarkupData
	ID              string
	IgnoreKeyboard  bool
	ShouldCrossFade bool
	ShowControls    bool
	ShowIndicators  bool
	DarkMode        bool
	DisableTouch    bool
	PauseOnHover    bool
	Interval        uint
	Autoplay        bool
	NoWrap          bool
	Items           []*CarouselItem
	ActiveItemIndex uint
}

func (c *Carousel) Render() vecty.ComponentOrHTML {
	if c.ID == "" {
		if c.ShowControls {
			panic("a carousel id is required to show carousel controls")
		}
		if c.ShowIndicators {
			panic("a carousel id is required to show carousel indicators")
		}
	}

	indicators := []vecty.MarkupOrChild{
		vecty.Markup(vecty.Class("carousel-indicators")),
	}

	items := []vecty.MarkupOrChild{
		vecty.Markup(vecty.Class("carousel-inner")),
	}

	for i, item := range c.Items {
		indicators = append(
			indicators,
			elem.Button(
				vecty.Markup(
					vecty.Type(vecty.TypeButton),
					vecty.Property{"aria-label": fmt.Sprintf("slide-%d", i)},
					vecty.Data{
						"bsTarget":  fmt.Sprintf("#%s", c.ID),
						"bsSlideTo": fmt.Sprintf("%d", i),
					},
				),
			),
		)

		item.isActive = uint(i) == c.ActiveItemIndex
		items = append(items, item)
	}

	pause := "hover"
	if !c.PauseOnHover {
		pause = "false"
	}

	interval := "5000"
	if c.Interval != 0 {
		interval = fmt.Sprintf("%d", c.Interval)
	}

	ride := "false"
	if c.Autoplay {
		ride = "carousel"
	}

	return elem.Div(
		vecty.Markup(
			c.MarkupData.Markup(),
			vecty.MarkupIf(c.ID != "", vecty.ID(c.ID)),
			vecty.Class("carousel", "slide"),
			vecty.ClassMap{
				"carousel-fade": c.ShouldCrossFade,
				"carousel-dark": c.DarkMode,
			},
			vecty.Data{
				"bsRide":     ride,
				"bsPause":    pause,
				"bsInterval": interval,
				"bsWrap":     fmt.Sprintf("%t", !c.NoWrap),
				"bsTouch":    fmt.Sprintf("%t", !c.DisableTouch),
				"keyboard":   fmt.Sprintf("%t", !c.IgnoreKeyboard),
			},
		),
		vecty.If(
			c.ShowControls,
			elem.Button(
				vecty.Markup(
					vecty.Class("carousel-control-prev"),
					vecty.Type(vecty.TypeButton),
					vecty.Data{"bsTarget": fmt.Sprintf("#%s", c.ID)},
				),
				elem.Span(
					vecty.Markup(
						vecty.Class("carousel-control-prev-icon"),
						vecty.Property{"aria-hidden": true},
					),
				),
				elem.Span(
					vecty.Markup(
						vecty.Class("visually-hidden"),
					),
					vecty.Text("Previous"),
				),
			),
			elem.Button(
				vecty.Markup(
					vecty.Class("carousel-control-next"),
					vecty.Type(vecty.TypeButton),
					vecty.Data{"bsTarget": fmt.Sprintf("#%s", c.ID)},
				),
				elem.Span(
					vecty.Markup(
						vecty.Class("carousel-control-prev-next"),
						vecty.Property{"aria-hidden": true},
					),
				),
				elem.Span(
					vecty.Markup(
						vecty.Class("visually-hidden"),
					),
					vecty.Text("Next"),
				),
			),
		),
		vecty.If(c.ShowIndicators, elem.Div(indicators...)),
		elem.Div(items...),
	)
}

type CarouselItem struct {
	vecty.Core
	internal.MarkupData
	isActive bool
	Interval uint
	Child    vecty.ComponentOrHTML
	Caption  *CarouselItemCaption
}

func (c *CarouselItem) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			c.MarkupData.Markup(),
			vecty.Class("carousel-item"),
			vecty.ClassMap{"active": c.isActive},
			vecty.MarkupIf(c.Interval != 0, vecty.Data{"bsInterval": fmt.Sprintf("%d", c.Interval)}),
		),
		c.Child,
		vecty.If(c.Caption != nil, c.Caption),
	)
}

type CarouselItemCaption struct {
	vecty.Core
	internal.MarkupData
	Child vecty.ComponentOrHTML
}

func (c *CarouselItemCaption) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			c.MarkupData.Markup(),
			vecty.Class("carousel-caption"),
		),
		c.Child,
	)
}
