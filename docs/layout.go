package docs

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"

	components "github.com/mvpmvh/vecty-bootstrap/v5"
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
					vecty.Text("Sidebar"),
				),
				elem.Section(
					vecty.Text("Main"),
				),
			),
		),
	)
}
