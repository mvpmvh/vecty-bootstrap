package components

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
)

type AccordionItem struct {
	vecty.Core
	Child vecty.ComponentOrHTML
}

type Accordion struct {
	vecty.Core
	Items []AccordionItem
}

func (a *Accordion) Render() vecty.ComponentOrHTML {
	return elem.Div()
}
