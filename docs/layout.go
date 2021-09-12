package docs

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"

	"github.com/mvpmvh/vecty-bootstrap/internal"
	components "github.com/mvpmvh/vecty-bootstrap/v5"
)

type Layout struct {
	vecty.Core
	Router *Router
	Main   func() vecty.ComponentOrHTML
}

func (l *Layout) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Div(
			&components.Container{
				Child: func() vecty.ComponentOrHTML {
					return vecty.Text("Header")
				},
			},
			elem.Div(
				vecty.Markup(
					vecty.Style{
						"display":               "grid",
						"grid-template-columns": "1fr 5fr",
						"grid-template-areas":   "'sidebar main'",
						"grid-gap":              "10px",
					},
				),
				elem.Section(
					&components.Accordion{
						ID:         "foo-accordion",
						AlwaysOpen: true,
						Items: []*components.AccordionItem{
							{
								Trigger: &components.AccordionTrigger{
									Container: func(c vecty.ComponentOrHTML) vecty.ComponentOrHTML {
										return elem.Heading2(
											vecty.Markup(
												vecty.Class("accordion-header"),
											),
											c,
										)
									},
									Button: &components.CollapseButton{
										Button: components.Button{
											Child: func() vecty.ComponentOrHTML {
												return vecty.Text("Components")
											},
										},
									},
								},
								Content: &components.Collapse{
									ID: "panelsStayOpen-collapseOne",
									Child: &components.Button{
										MarkupData: internal.MarkupData{
											Classes: []string{"accordion-body"},
											Data:    vecty.Data{"routerPath": "#alerts"},
											EventListeners: []*vecty.EventListener{
												event.Click(l.onRoute).PreventDefault(),
											},
										},
										Child: func() vecty.ComponentOrHTML {
											return vecty.Text("Alerts")
										},
									},
								},
							},
							{
								Trigger: &components.AccordionTrigger{
									Container: func(c vecty.ComponentOrHTML) vecty.ComponentOrHTML {
										return elem.Heading2(
											vecty.Markup(
												vecty.Class("accordion-header"),
											),
											c,
										)
									},
									Button: &components.CollapseButton{
										Button: components.Button{
											Child: func() vecty.ComponentOrHTML {
												return vecty.Text("Utilities")
											},
										},
									},
								},
								Content: &components.Collapse{
									ID: "panelsStayOpen-collapseTwo",
									Child: &components.Button{
										MarkupData: internal.MarkupData{
											Classes: []string{"accordion-body"},
											Data:    vecty.Data{"routerPath": "#background"},
											EventListeners: []*vecty.EventListener{
												event.Click(l.onRoute).PreventDefault(),
											},
										},
										Child: func() vecty.ComponentOrHTML {
											return vecty.Text("Background")
										},
									},
								},
							},
						},
					},
				),
				elem.Section(l.Main()),
			),
		),
	)
}

func (l *Layout) onRoute(e *vecty.Event) {
	path := e.Target.Get("dataset").Get("routerPath").String()
	l.Router.To(path)
}
