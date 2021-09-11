package docs

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	"github.com/mvpmvh/vecty-bootstrap/internal"
	components "github.com/mvpmvh/vecty-bootstrap/v5"
	"github.com/mvpmvh/vecty-bootstrap/v5/utilities"
)

type Layout struct {
	vecty.Core
}

func (l *Layout) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Div(
			&components.Container{
				Child: vecty.Text("Header"),
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
											Child: vecty.Text("Accordion Item #1"),
										},
									},
								},
								Content: &components.Collapse{
									ID: "panelsStayOpen-collapseOne",
									Child: elem.Div(
										vecty.Markup(
											vecty.Class("accordion-body"),
										),
										vecty.Tag("strong", vecty.Text("This is the first item's accordion body.")),
										vecty.Text("It is shown by default, until the collapse plugin adds the appropriate classes that we use to style each element. These classes control the overall appearance, as well as the showing and hiding via CSS transitions. You can modify any of this with custom CSS or overriding our default variables. It's also worth noting that just about any HTML can go within the <code>.accordion-body</code>, though the transition does limit overflow."),
									),
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
											Child: vecty.Text("Accordion Item #2"),
										},
									},
								},
								Content: &components.Collapse{
									ID: "panelsStayOpen-collapseTwo",
									Child: elem.Div(
										vecty.Markup(
											vecty.Class("accordion-body"),
										),
										vecty.Tag("strong", vecty.Text("This is the second item's accordion body.")),
										vecty.Text("It is not shown by default."),
									),
								},
							},
						},
					},
				),
				elem.Section(
					vecty.Text("Main"),
					&components.Alert{
						MarkupData:           internal.MarkupData{},
						Type:                 components.PrimaryAlert,
						DismissButton:        &components.Button{},
						ShouldAnimateDismiss: true,
						Child:                vecty.Text("A simple primary alert—check it out!"),
					},
					&components.Badge{
						MarkupData: internal.MarkupData{
							Classes: []string{utilities.BackgroundPrimary},
						},
						Child: vecty.Text("badge"),
					},
					&components.Badge{
						Child: vecty.Text("pill"),
						MarkupData: internal.MarkupData{
							Classes: []string{utilities.BackgroundSecondary, utilities.BorderRadiusRoundedPill},
						},
					},
				),
			),
		),
	)
}
