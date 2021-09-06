package components

import (
	"fmt"
	"strings"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/mvpmvh/vecty-bootstrap/internal"
)

// Slot is a function that receives a component so that it can be wrapped by a user-defined component
type Slot func(c vecty.ComponentOrHTML) vecty.ComponentOrHTML

// Accordion creates vertically collapsing content (see: https://getbootstrap.com/docs/5.1/components/accordion/)
type Accordion struct {
	vecty.Core
	internal.MarkupData
	ID      string
	Items   []*AccordionItem
	IsFlush bool
	// AlwaysOpen, if true, makes accordion items stay open when another item is opened (see: https://getbootstrap.com/docs/5.1/components/accordion/#always-open)
	AlwaysOpen bool
}

func (a *Accordion) Render() vecty.ComponentOrHTML {
	if a.AlwaysOpen && strings.TrimSpace(a.ID) == "" {
		panic("an accordion must have an ID if AlwaysOpen is true")
	}

	children := []vecty.MarkupOrChild{
		vecty.Markup(
			vecty.MarkupIf(a.ID != "", vecty.ID(a.ID)),
			vecty.Class("accordion"),
			vecty.ClassMap{
				"accordion-flush": a.IsFlush,
			},
		),
	}

	for _, item := range a.Items {
		item.parentAccordion = a
		children = append(children, item)
	}

	return elem.Div(children...)
}

// AccordionItem pairs a trigger (to hide/show content) with its content
type AccordionItem struct {
	vecty.Core
	internal.MarkupData
	Trigger         *AccordionTrigger
	Content         *Collapse
	parentAccordion *Accordion
}

func (a *AccordionItem) Render() vecty.ComponentOrHTML {
	a.Trigger.target = fmt.Sprintf("#%s", a.Content.ID)
	a.Content.Classes = append(a.Content.Classes, "accordion-collapse")
	if !a.parentAccordion.AlwaysOpen {
		a.Content.Data = append(a.Content.Data, vecty.Data{"bsParent": fmt.Sprintf("#%s", a.parentAccordion.ID)})
	}

	return elem.Div(
		vecty.Markup(
			vecty.Class("accordion-item"),
			a.MarkupData.Markup(),
		),
		a.Trigger,
		a.Content,
	)
}

// AccordionTrigger triggers the collapsing/showing of content. It is either a button or link.
type AccordionTrigger struct {
	vecty.Core
	// target is the dom selector (e.g. id or class) of the html element to be collapsed/shown
	target    string
	Container Slot
	Button    *CollapseButton
}

func (a *AccordionTrigger) Render() vecty.ComponentOrHTML {
	a.Button.Target = a.target
	a.Button.Classes = append(a.Button.Classes, "accordion-button")
	return a.Container(a.Button)
}
